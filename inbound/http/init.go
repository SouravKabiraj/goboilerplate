package http

import "go.uber.org/dig"

func Start(app *dig.Container) {
	initialize(app)
}
