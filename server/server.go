package server

import (
	"go.uber.org/dig"
	"sourav.kabiraj/goboilerplate/inbound/http"
	"sourav.kabiraj/goboilerplate/logic/todos"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	addHandlers(container)
	addManagers(container)
	return container
}

func addHandlers(container *dig.Container) {
	container.Provide(http.NewRouter)
}

func addManagers(container *dig.Container) {
	container.Provide(todos.NewManager)
}
