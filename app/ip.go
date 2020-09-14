package app

import (
	"github.com/HosseinZeinali/ip2loc/dto"
	"net"
)

/*func (ctx *Context) CreateIp(ip *model.Ip) {
	ctx.Database.CreateIp(ip)
}*/

func (ctx *Context) GetIpDetails(netIp net.IP) (*dto.IpDto, error) {
	if netIp.To4() != nil {
		ipInt := Ip2Int(netIp)

		ip, err := ctx.Database.GetIpv4RecordByIpv4(ctx.Database.DefaultIpv4Table, ipInt)
		ipDto := Ipv42IpDto(ip)
		if err != nil {
			return nil, err
		}
		return ipDto, nil
	} else {
		ipInt := Ipv6ToInt(netIp)
		ip, err := ctx.Database.GetIpv6RecordByIpv6(ctx.Database.DefaultIpv6Table, ipInt)
		ipDto := Ipv62IpDto(ip)
		if err != nil {
			return nil, err
		}
		return ipDto, nil
	}
}
