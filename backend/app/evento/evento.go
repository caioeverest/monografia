package evento

import (
	"encoding/json"
	"github.com/caioever/monografia/backend/adapters/mqtt"
	"github.com/caioever/monografia/backend/adapters/repositorioDeEventos"
	"github.com/caioever/monografia/backend/app/agendamentos"
)

type Evento struct {
	Agendamento			agendamentos.Agendamento	`json:"agendamento"`
	Commmand 			interface{}					`json:"commmand"`
}

func Builder(agendamento *agendamentos.Agendamento, comando interface{}) (*Evento, error){
	return &Evento{
		Agendamento: *agendamento,
		Commmand: comando,
	}, nil
}

var repository = repositorioDeEventos.Init()

func (e *Evento) Disparar(context *mqtt.Context, status string) error {
	if err := agendamentos.AtualizarStatusDoAgendamento(e.Agendamento.Id, status); err != nil {
		return err
	}
	if err := salvarEvento(e); err != nil {
		return err
	}
	if err := enviarComando(context, e); err != nil {
		return err
	}

	return nil
}

func TrazerTodosOsEventos() (interface{}, error) {
	return repository.FindAll()
}

func enviarComando(context *mqtt.Context, evento *Evento) error {
	msg, err := json.Marshal(evento.Commmand)
	if err != nil {
		return err
	}
	if err := context.Send(evento.Agendamento.Sala, msg); err != nil {
		return err
	}
	return nil
}

func salvarEvento(n *Evento) error {
	agendamento := &repositorioDeEventos.EventoEntity{
		Agendamento:	n.Agendamento,
		Commmand:		n.Commmand,
	}
	return repository.Save(agendamento)
}