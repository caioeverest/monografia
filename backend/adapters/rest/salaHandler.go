package rest

import (
	"github.com/caioever/monografia/backend/app/salas"
	"github.com/caioever/monografia/backend/infra/helper"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func (restBase *restBase) RecuperaTodasAsSala(c echo.Context) error {
	resultado := salas.TrazerTodasAsSalas()
	return c.JSON(http.StatusOK, resultado)
}

func (restBase *restBase) RecuperaTodasAsSalaDisponiveis(c echo.Context) error {
	inicio, errInicio := time.Parse("2006-01-02T15:04:05Z07:00", c.QueryParam("inicio"))
	fim, errFim := time.Parse("2006-01-02T15:04:05Z07:00",c.QueryParam("fim"))
	if errInicio != nil || errFim != nil {
		return c.JSON(http.StatusPreconditionFailed,
			helper.ResponseBuilder("Parametros de inicio e fim devem estar presentes e no formato correto",
				http.StatusPreconditionFailed))
	}
	resultado := salas.SalasDisponiveis(inicio, fim)
	return c.JSON(http.StatusOK, resultado)
}

func (restBase *restBase) RecuperaIdentificadoresDeTodasAsSala(c echo.Context) error {
	resultado, err := salas.TrazerIdentificadoresDeTodasAsSalas()
	if err != nil{
		return c.JSON(http.StatusServiceUnavailable,
			helper.ResponseBuilder(err,
				http.StatusServiceUnavailable))
	}
	return c.JSON(http.StatusOK, resultado)
}

func (restBase *restBase) CriarNovaSala(c echo.Context) error {
	body := &salas.Sala{}
	if err := c.Bind(body); err != nil{
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseBuilder(err.Error(),
				http.StatusInternalServerError))
	}
	err := salas.CriarNovaSala(body)
	if err != nil {
		return c.JSON(http.StatusConflict,
			helper.ResponseBuilder(err.Error(),
				http.StatusConflict))
	}
	return c.JSON(http.StatusCreated, nil)
}

func (restBase *restBase) RecuperaUmaSalaPeloId(c echo.Context) error {
	id := c.Param("id")
	log.Print(id)
	sala,_ := salas.RecuperaSalaPeloId(id)

	if sala == nil {
		return c.JSON(http.StatusNotFound,
			helper.ResponseBuilder("Sala não foi encontrato",
				http.StatusNotFound))
	}
	return c.JSON(http.StatusOK, sala)
}

func (restBase *restBase) RecuperaUmaSalaPeloIdentificador(c echo.Context) error {
	id := c.Param("identificador")
	log.Print(id)
	sala, err := salas.EncontrarSalaPorIdentificador(id)

	if err != nil {
		return c.JSON(http.StatusNotFound,
			helper.ResponseBuilder("Sala não foi encontrato",
				http.StatusNotFound))
	}
	return c.JSON(http.StatusOK, sala)
}

func (restBase *restBase) AtualizarSala(c echo.Context) error {
	id := c.Param("id")
	body := echo.Map{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseBuilder(err.Error(),
				http.StatusInternalServerError))
	}
	if err := salas.AtualizaSala(body, id); err != nil {
		return c.JSON(http.StatusNotFound,
			helper.ResponseBuilder(err.Error(),
				http.StatusNotFound))
	}
	return c.JSON(http.StatusAccepted,
		helper.ResponseBuilder("Atualizado com sucesso",
			http.StatusAccepted))
}

func (restBase *restBase) DeletarSala(c echo.Context) error {
	id := c.Param("id")
	err := salas.DeletaSala(bson.ObjectId(id).Hex())

	if err != nil{
		return c.JSON(http.StatusNotFound,
			helper.ResponseBuilder(err.Error(),
				http.StatusNotFound))
	}
	return c.JSON(http.StatusAccepted,
		helper.ResponseBuilder("Deletdo com sucesso",
			http.StatusAccepted))
}