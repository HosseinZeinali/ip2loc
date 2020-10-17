package http

import (
	"github.com/HosseinZeinali/ip2loc/app"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	App              *app.App
	TemplateRenderer *TemplateRenderer
	Router           *Router
	Api              *Api
}

func NewServer(app *app.App, renderer *TemplateRenderer, router *Router, api *Api) *Server {
	server := new(Server)
	server.TemplateRenderer = renderer
	server.App = app
	server.Router = router
	server.Api = api

	return server
}

func (s *Server) Render(name string, w http.ResponseWriter, view MainView) {
	buf, err := s.TemplateRenderer.Exec(name, view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(":8080", s.Router)
}

func (s *Server) Index() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		s.Render("index", w, MainView{})
	}
}

func (s *Server) InitRoutes() {
	s.Router.GET("/", s.Index())
}
