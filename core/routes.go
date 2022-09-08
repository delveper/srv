package core

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

// StdMux is simple std lib mux
type StdMux struct{ http.ServeMux }

// GrlMux is simple gorilla router
type GrlMux struct{ mux.Router }

// HttpRouter another one simple but fast router
type HttpRouter struct{ httprouter.Router }

// RegisterRoutes handle routes for each type of routers
func (sm *StdMux) RegisterRoutes(hdl ApiHandler) {
	sm.Handle(hdl.Endpoints.Root, Wrap(hdl, DefaultHandler, CheckMethod))
	sm.Handle(hdl.Endpoints.Index, Wrap(hdl, IndexHandler, CheckMethod))
	sm.Handle(hdl.Endpoints.Random, Wrap(hdl, CheckMethod, ValidateParams(hdl.Params...)))
}

func (gm *GrlMux) RegisterRoutes(hdl ApiHandler) {
	gm.Handle(hdl.Endpoints.Index, Wrap(hdl, IndexHandler)).Methods(http.MethodGet, http.MethodPost)
	gm.Handle(hdl.Endpoints.Root, Wrap(hdl, DefaultHandler)).Methods(http.MethodGet, http.MethodPost)
	gm.Handle(hdl.Endpoints.Random, Wrap(hdl, ValidateParams(hdl.Params...))).Methods(http.MethodGet, http.MethodPost)
}

func (hr *HttpRouter) RegisterRoutes(hdl ApiHandler) {
	hr.Handler(http.MethodGet, hdl.Endpoints.Root, Wrap(hdl, DefaultHandler))
	hr.Handler(http.MethodPost, hdl.Endpoints.Random, CheckMethod(hdl))
	hr.Handler(http.MethodGet, hdl.Endpoints.Random, Wrap(hdl, ValidateParams(hdl.Params...)))
	hr.Handler(http.MethodGet, hdl.Endpoints.Index, Wrap(hdl, IndexHandler))
}
