package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/andhikagama/os-api/di"
	"github.com/andhikagama/os-api/infrastructures"
	"github.com/andhikagama/os-api/migrations"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/dig"
)

func main() {
	container := di.BuildContainer()

	if err := runMigration(os.Args, container); err != nil {
		panic(err)
	}

	if err := runWebServer(container); err != nil {
		panic(err)
	}
}

func runMigration(cmd []string, container *dig.Container) error {
	if len(cmd) == 0 {
		return nil
	}

	if cmd[1] != "migrate" {
		return errors.New("unknown command")
	}

	if err := container.Invoke(func(mg *migrations.Migration) {
		mg.Up()
	}); err != nil {
		return err
	}

	fmt.Println("migrate success")
	os.Exit(0)
	return nil
}

func runWebServer(container *dig.Container) error {
	if err := container.Invoke(func(webServer *infrastructures.WebServer) {
		webServer.Serve()
	}); err != nil {
		return err
	}

	return nil
}
