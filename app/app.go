package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewApp() *App {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return &App{
		logger: logger,
	}
}

type App struct {
	router *mux.Router
	logger *zap.Logger
}

func (app *App) initRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", app.handleRoot)
	r.HandleFunc("/line-webhook", app.handleLine)
	app.router = r
}

func (app *App) Run(addr string) {
	app.initRoutes()
	srv := &http.Server{
		Handler: app.router,
		Addr:    addr,
	}

	srv.ListenAndServe()
}
