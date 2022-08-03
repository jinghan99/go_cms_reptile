package tools

import (
	"fmt"
	"testing"
)

//test 测试函数
func TestDnsListRecords(t *testing.T) {
	t.Logf("hello,world")
	records, err := DnsListRecords()
	if err != nil {
		panic(err)
	}
	fmt.Println(records)
}

func TestMyIp(t *testing.T) {
	ip, err := MyIp()
	if err != nil {
		panic(err)
	}
	fmt.Println(ip)
}

func TestMatchDomain(t *testing.T) {

	records, err := DnsListRecords()
	if err != nil {
		fmt.Println("nameSilo failed, err:", err)
		return
	}

	rrId, _ := MatchDomainRecordId(records)

	fmt.Println("nameSilo records success ", rrId)

	ip, _ := MyIp()
	fmt.Println("MyIp success ", ip.IP)

	err = UpdateDnsRecord(rrId, ip.IP)
	fmt.Println("nameSilo UpdateDnsRecord err ", err)
}
