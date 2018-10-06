package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (app *App) handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func (app *App) handleLine(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	app.logger.Debug(string(data))
}
