package main

import (
	"sourav.kabiraj/goboilerplate/inbound/http"
	"sourav.kabiraj/goboilerplate/server"
)

func main() {
	app := server.BuildContainer()
	http.Start(app)
}
