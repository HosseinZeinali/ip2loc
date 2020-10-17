package http

import (
	"encoding/json"
	"fmt"
	"github.com/HosseinZeinali/ip2loc/app"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
)

func (a *Api) IpHandler(ctx *app.Context) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ipString := p.ByName("ip")
		netIp := net.ParseIP(ipString)
		ip, err := ctx.GetIpDetails(netIp)
		data, err := json.Marshal(ip)

		_, err = w.Write(data)
		if err != nil {
			fmt.Println("error")
		}
	}
}
