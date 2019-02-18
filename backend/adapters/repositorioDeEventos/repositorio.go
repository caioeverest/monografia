package repositorioDeEventos

import "github.com/caioever/monografia/backend/adapters/database"

func Init() *database.Repository {
	r := database.Inject("eventos", EventoEntity{})
	return r
}
