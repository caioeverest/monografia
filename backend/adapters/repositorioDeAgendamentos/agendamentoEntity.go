package repositorioDeAgendamentos

import (
	"github.com/caioever/monografia/backend/adapters/database"
	"time"
)

type AgendamentoEntity struct {
	database.BaseDoc 	`bson:",inline"`
	Sala 				string
	Professor 			string
	Inicio 				time.Time
	Fim 				time.Time
	Status				string
}