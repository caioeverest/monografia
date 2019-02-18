package main

import (
	"github.com/caioever/monografia/backend/adapters/mqtt"
	"github.com/caioever/monografia/backend/app/evento"
	"github.com/caioever/monografia/backend/infra/config"
	"github.com/caioever/monografia/backend/adapters/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	httpConfig, mqttConf := config.CreateConfig()
	mqttContext := mqtt.ContextBuilder(mqttConf)
	go evento.IniciarPipeMonitoramento(mqttContext)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	router.Routes(e, mqttContext)

	e.Logger.Fatal(e.Start(httpConfig.Hostname+":"+httpConfig.Port))
}