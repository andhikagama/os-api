package main

import (
	"github.com/andhikagama/os-api/di"
	"github.com/andhikagama/os-api/infrastructures"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	container := di.BuildContainer()

	err := container.Invoke(func(webServer *infrastructures.WebServer) {
		webServer.Serve()
	})

	if err != nil {
		panic(err)
	}

}
