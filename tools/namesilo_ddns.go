// Package tools 工具库  域名 ddns 动态修改
//api :https://www.namesilo.com/api-reference#dns/dns-list-records
package tools

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"go_cms_reptile/models"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type MyIpMOdel struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	Cc      string `json:"cc"`
}
type Conf struct {
	NameSiloConf NameSiloConf `yaml:"nameSilo"`
}

// NameSiloConf 配置文件导入数据
type NameSiloConf struct {
	ApiKey   string `yaml:"apikey"`
	Domain   string `yaml:"domain"`
	DDnsHost string `yaml:"ddns_host"`
}

var yamlConf *Conf

// go 启动自动加载 init 函数
func init() {
	//     ./是你当前的工程目录，并不是该go文件所对应的目录
	yamlFile, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		panic("conf.yaml import error")
	}
	err = yaml.Unmarshal(yamlFile, &yamlConf)
	if err != nil {
		panic("yamlFile Unmarshal error")
	}

}

// MyIp 我的本地ip
func MyIp() (*MyIpMOdel, error) {
	httpUrl := "https://api.myip.com/"
	resp, err := http.Get(httpUrl)
	if err != nil {
		return nil, err
	}

	var myip *MyIpMOdel

	// body 正确响应 json  格式 {"ip":"118.112.111.89","country":"China","cc":"CN"}
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &myip); err != nil {
		return nil, err
	}

	return myip, nil
}

// DnsListRecords 获取 namesilo列出当前 DNS 记录
func DnsListRecords() (*models.NameSiloRecordModel, error) {

	httpUrl := "https://www.namesilo.com/api/dnsListRecords"

	req, _ := http.NewRequest(http.MethodGet, httpUrl, nil)
	req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 6.0.1; VTR-AL00 Build/V417IR)")

	//设置查询参数
	query := req.URL.Query()
	query.Add("domain", yamlConf.NameSiloConf.Domain)
	query.Add("key", yamlConf.NameSiloConf.ApiKey)
	query.Add("type", "xml")
	query.Add("version", "1")

	req.URL.RawQuery = query.Encode()
	//发起 请求 响应
	resp, _ := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	// body 正确响应 xml 格式
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var record *models.NameSiloRecordModel

	//判断是否转换失败
	if err = xml.Unmarshal(body, &record); err != nil {
		return nil, err
	}
	return record, nil
}

// MatchDomain 匹配 地址  ddns 获取 rrid
func MatchDomainRecordId(record *models.NameSiloRecordModel) (string, error) {
	resourceRecord := record.Reply.ResourceRecord

	for index, value := range resourceRecord {
		fmt.Println(index, value)
		// 匹配成功 ddns 需要的 域名 返回
		if value.Host.Text == yamlConf.NameSiloConf.DDnsHost+"."+yamlConf.NameSiloConf.Domain {
			return value.RecordID.Text, nil
		}
	}
	return "", errors.New("no match ddns domain")
}

// UpdateDnsRecord Update an existing DNS resource record.
//https://www.namesilo.com/api/dnsUpdateRecord?version=1
//&type=xml
//&key=12345
//&domain=namesilo.com
//&rrid=1a2b3 rrid：资源记录的唯一 ID。您可以使用 dnsListRecords 获取此值。
//&rrhost=test
//&rrvalue=55.55.55.55
//&rrttl=7207
func UpdateDnsRecord(rrId, updateValue string) error {
	httpUrl := "https://www.namesilo.com/api/dnsUpdateRecord?version=1&type=xml"

	req, _ := http.NewRequest(http.MethodGet, httpUrl, nil)
	req.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 6.0.1; VTR-AL00 Build/V417IR)")

	//设置查询参数
	query := req.URL.Query()
	query.Add("version", "1")
	query.Add("type", "xml")
	query.Add("key", yamlConf.NameSiloConf.ApiKey)
	query.Add("domain", yamlConf.NameSiloConf.Domain)
	query.Add("rrid", rrId)
	query.Add("rrhost", yamlConf.NameSiloConf.DDnsHost)
	query.Add("rrvalue", updateValue)
	query.Add("rrttl", "3603")

	req.URL.RawQuery = query.Encode()
	//发起 请求 响应
	resp, _ := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// body 正确响应 xml 格式
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var updateResp *models.UpDateDnsRecordRespModel
	if err := xml.Unmarshal(body, &updateResp); err != nil {
		return err
	}

	if "success" == updateResp.Reply.Detail.Text {
		log.Printf("name-silo UpdateDnsRecord ,rrhost：%v , value：%v \n", yamlConf.NameSiloConf.DDnsHost, updateValue)
		return nil
	}

	return errors.New("UpdateDnsRecord error")
}
