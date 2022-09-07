package core

import (
	"fmt"
	"net/http"

	"github.com/delveper/srv/cfg"
)

// Muxer designed to make code DRY
// while running different types of configurations
type Muxer interface {
	RegisterRoutes(ApiHandler)
	http.Handler
}

// Run is main function
// that rise a server with given options
func Run(mux Muxer) error {
	config, err := cfg.Load()
	if err != nil {
		return fmt.Errorf("error loading congfig: %w", err)
	}

	hdl := ApiHandler{config}

	srv := &http.Server{
		Handler:      mux,
		Addr:         hdl.HTTP.Host + ":" + hdl.HTTP.Port,
		ReadTimeout:  hdl.HTTP.ReadTimeout,
		WriteTimeout: hdl.HTTP.WriteTimeout,
		IdleTimeout:  hdl.HTTP.IdleTimeout,
	}

	mux.RegisterRoutes(hdl)

	if err = srv.ListenAndServe(); err != nil {
		return fmt.Errorf("error loading the service: %w", err)
	}

	return nil
}
