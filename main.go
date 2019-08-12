package main

import (
	"github.com/senowijayanto/gorm-mux/app"
	"github.com/senowijayanto/gorm-mux/app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
