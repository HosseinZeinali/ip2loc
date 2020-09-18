package http

import (
	"github.com/HosseinZeinali/ip2loc/app"
	"github.com/gorilla/mux"
	"net/http"
)

type API struct {
	Config *Config
	App    *app.App
}

func New(a *app.App) (api *API, err error) {
	api = &API{App: a}
	api.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}
	return api, nil
}

func (a *API) Init(r *mux.Router) {
	router := r.PathPrefix("/").Subrouter()
	router.Handle("/ip/{ip}", a.handler(a.GetIpDetails)).Methods("GET")
}

func (a *API) handler(f func(*app.Context, http.ResponseWriter, *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := a.App.NewContext()
		if err := f(ctx, w, r); err != nil {
			ctx.Logger.ActionError("handling http request.")

			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	})
}
