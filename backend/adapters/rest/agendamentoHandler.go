package rest

import (
	"github.com/caioever/monografia/backend/app/agendamentos"
	"github.com/caioever/monografia/backend/infra/helper"
	"github.com/labstack/echo"
	"net/http"
)

func (restBase *restBase) RecuperaTodosOsAgendamentos(c echo.Context) error {
	resultados, err := agendamentos.TrazerTodosAgendamentos()
	if err != nil{
		return c.JSON(http.StatusServiceUnavailable,
			helper.ResponseBuilder(err,
				http.StatusServiceUnavailable))
	}
	return c.JSON(http.StatusOK, resultados)
}

func (restBase *restBase) CriarNovoAgendamento(c echo.Context) error {
	body := &agendamentos.Agendamento{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseBuilder(err.Error(),
			http.StatusInternalServerError))
	}
	if err := agendamentos.CriarNovoAgendamento(body); err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseBuilder(err.Error(),
				http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, body)
}

func (restBase *restBase) RecuperaTodosUmAgendamentoPeloId(c echo.Context) error {
	id := c.Param("id")
	agendamento, err :=agendamentos.RecuperaAgendamentoPeloId(id)
	if err != nil{
		return c.JSON(http.StatusNotFound,
			helper.ResponseBuilder("Agendamento não encontrado",
			http.StatusNotFound))
	}
	return c.JSON(http.StatusOK, agendamento)
}

func (restBase *restBase) AtualizarAgendamento(c echo.Context) error {
	id := c.Param("id")
	body := echo.Map{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseBuilder(err.Error(),
				http.StatusInternalServerError))
	}
	if err := agendamentos.AtualizaAgendamento(body, id); err != nil {
		return c.JSON(http.StatusConflict,
			helper.ResponseBuilder(err.Error(),
				http.StatusConflict))
	}
	return c.JSON(http.StatusAccepted,
		helper.ResponseBuilder("Atualizado com sucesso",
			http.StatusAccepted))
}

func (restBase *restBase) DeletarAgendamento(c echo.Context) error {
	id := c.Param("id")
	err := agendamentos.DeletaAgendamento(id)

	if err != nil{
		return c.JSON(http.StatusNotFound,
			helper.ResponseBuilder("Agendamento não encontrado",
				http.StatusNotFound))
	}
	return c.JSON(http.StatusAccepted,
		helper.ResponseBuilder("Deletdo com sucesso",
			http.StatusAccepted))
}