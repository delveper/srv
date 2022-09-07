package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

// StdMux is simple std lib mux
type StdMux struct{ http.ServeMux }

// GrlMux is simple gorilla router
type GrlMux struct{ mux.Router }

// RegisterRoutes handle routes for StdMux
func (sm *StdMux) RegisterRoutes(hdl ApiHandler) {
	sm.Handle(hdl.Endpoints.Root, Wrap(hdl, DefaultHandler, CheckMethod))
	sm.Handle(hdl.Endpoints.Index, Wrap(hdl, IndexHandler, CheckMethod))
	sm.Handle(hdl.Endpoints.Random, Wrap(hdl, CheckMethod, ValidateParams(hdl.Params...)))
}

// RegisterRoutes is what it is
func (gm *GrlMux) RegisterRoutes(hdl ApiHandler) {
	gm.Handle(hdl.Endpoints.Index, Wrap(hdl, IndexHandler)).Methods(http.MethodGet, http.MethodPost)
	gm.Handle(hdl.Endpoints.Root, Wrap(hdl, DefaultHandler)).Methods(http.MethodGet, http.MethodPost)
	gm.Handle(hdl.Endpoints.Random, Wrap(hdl, ValidateParams(hdl.Params...))).Methods(http.MethodGet, http.MethodPost)
}
