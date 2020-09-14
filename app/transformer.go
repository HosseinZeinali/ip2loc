package app

import (
	"encoding/binary"
	"github.com/HosseinZeinali/ip2loc/dto"
	"github.com/HosseinZeinali/ip2loc/model"
	"math/big"
	"net"
	"strconv"
	"strings"
	"time"
)

func IpDto2Ipv4(dto *dto.IpDto) *model.Ipv4 {
	ipStart := net.ParseIP(dto.IpStart)
	ipStartInt := Ip2Int(ipStart)
	ipEnd := net.ParseIP(dto.IpEnd)
	ipEndInt := Ip2Int(ipEnd)
	ip := model.NewIpv4(ipStartInt, ipEndInt, dto.IpCount, dto.Date, dto.Status, dto.Cc)

	return ip
}

func IpDto2Ipv6(dto *dto.IpDto) *model.Ipv6 {
	ipStart := net.ParseIP(dto.IpStart)
	ipStartInt := Ipv6ToInt(ipStart)
	ipEnd := net.ParseIP(dto.IpEnd)
	ipEndInt := Ipv6ToInt(ipEnd)
	ip := model.NewIpv6(ipStartInt, ipEndInt, dto.IpCount, dto.Date, dto.Status, dto.Cc)

	return ip
}

func Ip2Int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func Int2Ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func Ipv6ToInt(IPv6Addr net.IP) *big.Int {
	IPv6Int := big.NewInt(0)
	IPv6Int.SetBytes(IPv6Addr)
	return IPv6Int
}

func IntToIpv6(intipv6 *big.Int) net.IP {
	ip := intipv6.Bytes()
	var a net.IP = ip
	return a
}

func Nic2Ipv4Dto(nic string) *dto.IpDto {
	nicRecord := strings.Split(nic, "|")
	ipCount, _ := strconv.Atoi(nicRecord[4])
	ipStart := net.ParseIP(nicRecord[3])
	ipStartInt := Ip2Int(ipStart)
	ipEndInt := ipStartInt + uint32(ipCount)
	ipEnd := Int2Ip(ipEndInt)
	ip := dto.NewIpDto(nicRecord[2], nicRecord[3], ipEnd.String(), ipCount, time.Now(), nicRecord[6], nicRecord[1])

	return ip
}

func Nic2Ipv6Dto(nic string) *dto.IpDto {
	nicRecord := strings.Split(nic, "|")
	ipCount, _ := strconv.Atoi(nicRecord[4])
	ipStart := net.ParseIP(nicRecord[3])
	ipStartInt := Ipv6ToInt(ipStart)
	ipCount2 := big.NewInt(int64(ipCount))
	ipEndInt := big.NewInt(0).Add(ipStartInt, ipCount2)
	ipEnd := IntToIpv6(ipEndInt)
	ip := dto.NewIpDto(nicRecord[2], nicRecord[3], ipEnd.String(), ipCount, time.Now(), nicRecord[6], nicRecord[1])

	return ip
}

func Ipv62IpDto(ipv6 *model.Ipv6) *dto.IpDto {
	ipStartInt := new(big.Int)
	ipStartInt.SetString(ipv6.IpStart, 10)
	ipStart := IntToIpv6(ipStartInt)
	ipEndInt := new(big.Int)
	ipEndInt.SetString(ipv6.IpEnd, 10)
	ipEnd := IntToIpv6(ipStartInt)
	ip := dto.NewIpDto("ipv6", ipStart.String(), ipEnd.String(), ipv6.IpCount, ipv6.Date, ipv6.Status, ipv6.Cc)

	return ip
}

func Ipv42IpDto(ipv4 *model.Ipv4) *dto.IpDto {
	ipStart := Int2Ip(ipv4.IpStart)
	ipEnd := Int2Ip(ipv4.IpEnd)
	ip := dto.NewIpDto("ipv6", ipStart.String(), ipEnd.String(), ipv4.IpCount, ipv4.Date, ipv4.Status, ipv4.Cc)

	return ip
}
