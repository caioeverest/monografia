package repositorioDeEventos

import (
	"github.com/caioever/monografia/backend/adapters/database"
	"github.com/caioever/monografia/backend/app/agendamentos"
)

type EventoEntity struct {
	database.BaseDoc 	`bson:",inline"`
	Agendamento			agendamentos.Agendamento
	Commmand 			interface{}
}

