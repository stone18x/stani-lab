package main

import (
	"github.com/gorilla/mux"
)

func main() {
	app := App{Router: mux.NewRouter()}
	app.Run(8080)
}
