package router

import (
	"github.com/caioever/monografia/backend/adapters/mqtt"
	"github.com/caioever/monografia/backend/adapters/rest"
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo, mqttContext *mqtt.Context) {

	handler := rest.InjectConfig(mqttContext)

	//Health check
	e.GET("/health", handler.Health)

	//Agendamentos
	e.GET("/agenda", handler.RecuperaTodosOsAgendamentos)
	e.POST("/agendar", handler.CriarNovoAgendamento)
	e.GET("/agendamento/:id", handler.RecuperaTodosUmAgendamentoPeloId)
	e.PUT("/agendamento/:id", handler.AtualizarAgendamento)
	e.DELETE("/agendamento/:id", handler.DeletarAgendamento)

	//Salas
	e.GET("/salas", handler.RecuperaTodasAsSala)
	e.GET("/salas/disponiveis", handler.RecuperaTodasAsSalaDisponiveis)
	e.GET("/salas/identificadores", handler.RecuperaIdentificadoresDeTodasAsSala)
	e.POST("/sala", handler.CriarNovaSala)
	e.GET("/sala/:id", handler.RecuperaUmaSalaPeloId)
	e.GET("/sala/identificador/:identificador", handler.RecuperaUmaSalaPeloIdentificador)
	e.PUT("/sala/:id", handler.AtualizarSala)
	e.DELETE("/sala/:id", handler.DeletarSala)

	//Comando
	e.GET("/evento", handler.RecuperaTodosOsEventos)
}
