package http

import (
	"encoding/json"
	"github.com/HosseinZeinali/ip2loc/app"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

//
func (a *API) GetIpDetails(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	netIp := net.ParseIP(getStringIpFromRequest(r))
	ip, err := ctx.GetIpDetails(netIp)
	data, err := json.Marshal(ip)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func getStringIpFromRequest(r *http.Request) string {
	vars := mux.Vars(r)
	stringIp := vars["ip"]

	return stringIp
}
