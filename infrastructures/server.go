package infrastructures

import (
	"fmt"

	"github.com/andhikagama/os-api/infrastructures/rest"
	"github.com/andhikagama/os-api/shared"
)

type WebServer struct {
	resource shared.Resource
	routes   rest.Routes
}

func (ws *WebServer) Serve() {

	// Set validator
	ws.resource.Echo.Validator = ws.resource.Validator

	v1 := ws.resource.Echo.Group(`/v1`)
	// v1.Use(ws.middleware.SetCORS())

	ws.routes.Register(v1)

	ws.resource.Echo.Start(fmt.Sprintf(`:%v`, ws.resource.Config.ServerPort))
}

func New(resource shared.Resource,
	routes rest.Routes,
) *WebServer {
	return &WebServer{
		resource: resource,
		routes:   routes,
	}
}
