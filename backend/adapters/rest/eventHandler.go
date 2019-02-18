package rest

import (
	"github.com/caioever/monografia/backend/app/evento"
	"github.com/caioever/monografia/backend/infra/helper"
	"github.com/labstack/echo"
	"net/http"
)

func (restBase *restBase) RecuperaTodosOsEventos(c echo.Context) error {
	resultados, err := evento.TrazerTodosOsEventos()
	if err != nil{
		return c.JSON(http.StatusServiceUnavailable,
			helper.ResponseBuilder(err,
				http.StatusServiceUnavailable))
	}
	return c.JSON(http.StatusOK, resultados)
}

//func (restBase *restBase) DispararEvento(c echo.Context) error {
//	id := c.Param("id")
//	command := lampadas.MontaComandoParaLampada(true)
//	event, err := evento.Builder(id, command)
//	if err != nil {
//		return c.JSON(http.StatusServiceUnavailable,
//			helper.ResponseBuilder(err.Error(),
//				http.StatusServiceUnavailable))
//	}
//	if err := event.Disparar(restBase.Context); err != nil {
//		return c.JSON(http.StatusServiceUnavailable,
//			helper.ResponseBuilder("Não foi possivel enviar o evento",
//				http.StatusServiceUnavailable))
//	}
//
//	resposta := make(map[string]interface{})
//	resposta["message"] = "Eventos enviado com sucesso"
//	resposta["evento"] = event
//
//	return c.JSON(http.StatusOK,
//		helper.ResponseBuilder( resposta, http.StatusOK))
//}