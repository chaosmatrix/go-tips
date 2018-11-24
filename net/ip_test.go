package lib

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	addrs := []string{
		"0.0.0.0",
		"0.0.0.110",
		"1.1.1.1",
		"1.2.3.4",
		"10.0.0.1",
		"10.10.10.10",
		"127.0.0.1",
		"127.0.0.200",
		"127.0.100.200",
		"127.100.100.200",
		"192.168.0.1",
		"200.100.10.0",
		"255.255.255.255",
	}

	for _, addr := range addrs {
		dec := Ipv4AddrToDec(addr)
		dec2 := Ipv4AddrToDec2(addr)
		_addr := DecToIpv4Addr(dec)
		fmt.Printf("%s -> %d -> %d -> %s\n", addr, dec, dec2, _addr)
	}
}

func BenchmarkIpToDec(b *testing.B) {
	addr := "255.255.255.255"
	for i := 0; i < b.N; i++ {
		Ipv4AddrToDec(addr)
	}
}
func BenchmarkIpToDec2(b *testing.B) {
	addr := "255.255.255.255"
	for i := 0; i < b.N; i++ {
		Ipv4AddrToDec2(addr)
	}
}
func BenchmarkDecToIp(b *testing.B) {
	var addrInt int64 = 4294967295
	for i := 0; i < b.N; i++ {
		DecToIpv4Addr(addrInt)
	}
}
