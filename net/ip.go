package lib

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// IP Address to Decimal
// 183 ns/op
func Ipv4AddrToDec(addr string) (dec int64) {
	addrs := strings.Split(addr, ".")
	if len(addrs) != 4 {
		return -1
	}
	for i := 3; i >= 0; i-- {
		if addrInt, err := strconv.ParseInt(addrs[3-i], 10, 64); err == nil {
			dec |= addrInt << uint(i*8)
		} else {
			return -1
		}
	}
	return
}

// 40 ns/op
func Ipv4AddrToDec2(addr string) (dec int64) {
	if len(addr) > 15 {
		return -1
	}
	prevIdx, idx, num := 0, 0, 0
	var shift uint = 3
	for ; idx < len(addr); idx++ {
		if addr[idx] < '0' || addr[idx] > '9' {
			if addr[idx] != '.' {
				return -1
			} else {
				for i := idx - 1; i >= prevIdx; i-- {
					num = num + int(addr[i]-'0')*int(math.Pow10(idx-1-i))
				}
				dec |= int64(num) << uint(shift*8)
				shift--
				prevIdx = idx + 1
				num = 0
			}
		}
	}
	for i := idx - 1; i >= prevIdx; i-- {
		num = num + int(addr[i]-'0')*int(math.Pow10(idx-1-i))
	}
	dec |= int64(num)
	return
}

/*
func DecToIpv4Addr(addrInt int64) (addr string) {
    var builder strings.Builder
    builder.Grow(7)
    for i := 3; i > 0; i-- {
        builder.WriteString(strconv.Itoa(int((addrInt >> uint(i*8)) & 0xFF)))
        builder.WriteString(".")
    }
    builder.WriteString(strconv.Itoa(int((addrInt >> 0) & 0xFF)))
    return builder.String()
}
*/

// 380 ns/op
func DecToIpv4Addr1(addrInt int64) string {
	return fmt.Sprintf(
		"%d.%d.%d.%d",
		(addrInt>>24)&0xFF,
		(addrInt>>16)&0xFF,
		(addrInt>>8)&0xFF,
		addrInt&0xFF)
}

// 400 ns/op
func DecToIpv4Addr2(addrInt int64) string {
	bs := make([]byte, 7, 15)
	buf := bytes.NewBuffer(bs)
	buf.WriteString(strconv.FormatInt((addrInt>>24)&0xFF, 10))
	buf.WriteString(".")
	buf.WriteString(strconv.FormatInt((addrInt>>16)&0xFF, 10))
	buf.WriteString(".")
	buf.WriteString(strconv.FormatInt((addrInt>>8)&0xFF, 10))
	buf.WriteString(".")
	//buf.WriteRune('.')
	buf.WriteString(strconv.FormatInt(addrInt&0xFF, 10))
	return buf.String()
}

// 55.0 ns/op
func DecToIpv4Addr(addrInt int64) string {
	bs := make([]byte, 17)

	idx := 0
	var i uint = 0
	for ; i < 4; i++ {
		dec := int(addrInt >> (8 * i) & 0xFF)
		for dec/10 != 0 {
			bs[16-idx] = byte(dec%10) + '0'
			dec = dec / 10
			idx++
		}
		bs[16-idx] = byte(dec%10) + '0'
		idx++
		bs[16-idx] = '.'
		idx++
	}
	return string(bs[18-idx:])
}
