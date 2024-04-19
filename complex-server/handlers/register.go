package handlers

import (
	"net/http"

	"github.com/onedaydev/mygolang/complex-server/config"
)

func Register(mux *http.ServeMux, conf config.AppConfig) {
	mux.Handle(
		"/healthz",
		&app{conf: conf, handler: healthCheckHandler},
	)
	mux.Handle(
		"/api",
		&app{conf: conf, handler: apiHandler},
	)
	mux.Handle(
		"/panic",
		&app{conf: conf, handler: panicHandler},
	)
}
