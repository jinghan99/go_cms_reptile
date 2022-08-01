// Package tools 工具库  域名 ddns 动态修改
//api :https://www.namesilo.com/api-reference#dns/dns-list-records
package tools

import (
	"encoding/json"
	"encoding/xml"
	"go_cms_reptile/models"
	"io"
	"io/ioutil"
	"net/http"
)

type MyIpMOdel struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	Cc      string `json:"cc"`
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

// DnsListRecords 获取 namesilo查看与域关联的所有当前 DNS 记录
func DnsListRecords(domain, apikey string) (*models.NamesiloRecordModel, error) {

	http_url := "https://www.namesilo.com/api/dnsListRecords"

	req, _ := http.NewRequest(http.MethodGet, http_url, nil)

	//设置查询参数
	query := req.URL.Query()
	query.Add("domain", domain)
	query.Add("key", apikey)
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
	var record *models.NamesiloRecordModel

	//判断是否转换失败
	if err = xml.Unmarshal(body, &record); err != nil {
		return nil, err
	}
	return record, nil
}
