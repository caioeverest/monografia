package evento

import (
	"log"
	"time"

	"github.com/caioever/monografia/backend/adapters/mqtt"
	"github.com/caioever/monografia/backend/app/agendamentos"
	"github.com/caioever/monografia/backend/app/comando/lampadas"
	"github.com/jasonlvhit/gocron"
)

//go mqttContext.Listen("orquestrador")

func IniciarPipeMonitoramento(context *mqtt.Context) {
	log.Print("Iniciando monitores")

	gocron.Every(1).Minute().Do(monitorDeInicio, context)
	gocron.Every(1).Minute().Do(monitorDeFechamento, context)

	<-gocron.Start()
}

func monitorDeInicio(context *mqtt.Context) {
	query := make(map[string]interface{})
	param := make(map[string]time.Time)

	param["$gte"] = time.Now()
	//param["$gte"] 	= time.Now().Add(-time.Minute)
	query["inicio"] = param
	query["status"] = agendamentos.EM_ESPERA

	result := agendamentos.Monitorar(query)
	if result == nil {
		return
	}

	log.Printf("Foram encontrados [%d] agendamentos para serem iniciados", len(result))
	for _, item := range result {
		comando := lampadas.MontaComandoParaLampada(true)
		event, err := Builder(agendamentos.EntitySimplifier(item), comando)
		if err != nil {
			log.Printf("Ocorreu um erro ao processar o agendamento %s Erro: %s", item.Id.Hex(), err)
		}
		err = event.Disparar(context, agendamentos.INICIADO)
		if err != nil {
			log.Printf("Ocorreu um erro ao disparar o evento do agendamento %s Erro: %s", item.Id.Hex(), err)
		}
	}
}

func monitorDeFechamento(context *mqtt.Context) {
	query := make(map[string]interface{})
	param := make(map[string]time.Time)

	param["$lt"] = time.Now()
	//param["$gte"] 	= time.Now().Add(-time.Minute)
	query["fim"] = param
	query["status"] = agendamentos.INICIADO

	result := agendamentos.Monitorar(query)
	if result == nil {
		return
	}

	log.Printf("Foram encontrados [%d] agendamentos para serem finalizados", len(result))
	for _, item := range result {
		comando := lampadas.MontaComandoParaLampada(false)
		event, err := Builder(agendamentos.EntitySimplifier(item), comando)
		if err != nil {
			log.Printf("Ocorreu um erro ao processar o agendamento %s Erro: %s", item.Id.Hex(), err)
		}
		err = event.Disparar(context, agendamentos.FINALIZADO)
		if err != nil {
			log.Printf("Ocorreu um erro ao disparar o evento do agendamento %s Erro: %s", item.Id.Hex(), err)
		}
	}
}
