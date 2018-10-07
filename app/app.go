package app

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gorilla/mux"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.uber.org/zap"
)

func NewApp(config *AppConfig) *App {
	logger, err := zap.NewDevelopment()

	if err != nil {
		panic(err)
	}

	bot, err := linebot.New(config.LineChannelSecret, config.LineAccessToken)
	if err != nil {
		panic(err)
	}

	return &App{
		logger: logger,
		bot:    bot,
	}
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		LineChannelSecret: viper.GetString("line.channel-secret"),
		LineAccessToken:   viper.GetString("line.access-token"),
		Address:           fmt.Sprint(viper.GetString("app.host"), ":", viper.GetInt("app.port")),
		Name:              viper.GetString("app.bot-name"),
	}
}

type AppConfig struct {
	LineChannelSecret string
	LineAccessToken   string
	Address           string
	Name              string
}

type App struct {
	router *mux.Router
	logger *zap.Logger
	bot    *linebot.Client
	// bot linebot.
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

	err := srv.ListenAndServe()
	if err != nil {
		app.logger.Fatal("Fail to Listen and serve..."+err.Error(), zap.Any("err stack", err))
	}
}
