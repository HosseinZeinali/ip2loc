package cmd

import (
	"github.com/HosseinZeinali/ip2loc/app"
	http "github.com/HosseinZeinali/ip2loc/http"
	"github.com/HosseinZeinali/ip2loc/model"
	"math/rand"
	"time"
)

func Serve() {
	ticker := time.NewTicker(6 * time.Hour)
	app, _ := app.New()
	ctx := app.NewContext()

	if !ctx.Database.DoesTableExist("tables") {
		ctx.Database.CreateTableTable()
	}

	go func() {
		for ; true; <-ticker.C {
			updateIpsIfNeeded(ctx)
		}
	}()

	renderer := http.NewTemplateRenderer()
	router := http.NewRouter()
	api := http.NewApi(router, app)
	server := http.NewServer(app, renderer, router, api)
	server.Api.InitRoutes()
	server.InitRoutes()
	server.ListenAndServe()
	//api1, _ := apiHttp.New(app)
	//router := mux.NewRouter()
	//cors := handlers.CORS(
	//	handlers.AllowedOrigins([]string{"*"}),
	//	handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS"}),
	//	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	//)
	//s := &http.Server{
	//	Addr:        ":8081",
	//	Handler:     cors(router),
	//	ReadTimeout: 2 * time.Minute,
	//}
	//api1.Init(router.PathPrefix("/api").Subrouter())
	//if err := s.ListenAndServe(); err != http.ErrServerClosed {
	//	logrus.Error(err)
	//}
}

func updateIpsIfNeeded(ctx *app.Context) {
	isChanged, _ := ctx.Nic.CheckForChange()
	if isChanged {
		ctx.Logger.ActionInfo("The database should be updated")
	} else {
		ctx.Logger.ActionInfo("The database is already up-to-date.")
	}
	if isChanged {
		ipv4Table := "ipv4s_" + RandStringRunes(8)
		ipv6Table := "ipv6s_" + RandStringRunes(8)
		ctx.Logger.ActionInfo("donwloding nic data")
		ctx.Nic.DownloadNicData()
		ctx.Database.CreateIpv4Table(ipv4Table)
		ctx.Database.CreateIpv6Table(ipv6Table)
		reader, _ := ctx.Nic.GetNicRecords()
		i := 1
		j := 1
		var ipv4s []*model.Ipv4
		var ipv6s []*model.Ipv6
		for ip := range reader {
			if ip.IpType == "ipv4" {
				ipv4 := app.IpDto2Ipv4(ip)
				ipv4s = append(ipv4s, ipv4)
				if i%100 == 0 {
					err := ctx.Database.CreateBatchIpv4s(ipv4Table, ipv4s)
					if err != nil {
						ctx.Logger.ActionError("inserting 100 ipv4 records")
					} else {
						ctx.Logger.ActionInfo("Successfully inserted 100 ipv4 records")
					}
					ipv4s = nil
				}
				i++
			}
			if ip.IpType == "ipv6" {
				ipv6 := app.IpDto2Ipv6(ip)
				ipv6s = append(ipv6s, ipv6)
				if j%100 == 0 {
					err := ctx.Database.CreateBatchIpv6s(ipv6Table, ipv6s)
					if err != nil {
						ctx.Logger.ActionError("inserting 100 ipv6 records")
					} else {
						ctx.Logger.ActionInfo("Successfully inserted 100 ipv6 records")
					}
					ipv6s = nil
				}
				j++
			}
		}
		ctx.Database.CreateBatchIpv4s(ipv4Table, ipv4s)
		ctx.Database.CreateBatchIpv6s(ipv6Table, ipv6s)
		ipv4T := model.NewTable(ipv6Table, "", time.Now(), "ipv6")
		ipv6T := model.NewTable(ipv4Table, "", time.Now(), "ipv4")
		ctx.Database.CreateTable(ipv4T)
		ctx.Database.CreateTable(ipv6T)
		ctx.Database.DefaultIpv6Table = ipv6Table
		ctx.Database.DefaultIpv4Table = ipv4Table
		ctx.Logger.ActionInfo("Successfully updated database")
	}
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
