package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
}

func (app *App) Run(addr string) {
	srv := &http.Server{
		Handler: app.router,
		Addr:    addr,
	}

	srv.ListenAndServe()
}
