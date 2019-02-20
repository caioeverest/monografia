package repositorioDeAgendamentos

import (
	"github.com/caioever/monografia/backend/adapters/database"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type RepositorioDeAgendamentos struct {
	*database.Repository
}

func Init() *RepositorioDeAgendamentos {
	r := database.Inject("agendamentos", AgendamentoEntity{})
	return &RepositorioDeAgendamentos{r}
}

func (repository *RepositorioDeAgendamentos) Monitorar(pipeQuery bson.M) []AgendamentoEntity {
	var list []AgendamentoEntity
	if err := repository.Collection().
		Find(pipeQuery).All(&list); err != nil {
		return nil
	}
		return list
}

func (repository *RepositorioDeAgendamentos) RetornarAgendamentoEntre(inicio, fim time.Time) []AgendamentoEntity {
	query := make(map[string]interface{})

	query["inicio"] = map[string]time.Time {
		"$lte": fim,
	}
	query["fim"] = map[string]time.Time {
		"$gte": inicio,
	}
	query["status"] = map[string]map[string]string {
		"$not": {
			"$gte": "FINALIZADO",
		},
	}

	var list []AgendamentoEntity
	if err := repository.Collection().
		Find(query).All(&list); err != nil {
		return nil
	}
	return list
}

func (repository *RepositorioDeAgendamentos) RecuperaAtivosEFormataAgendamentos() ([]AgendamentoEntity) {
	var list []AgendamentoEntity
	if err := repository.Collection().
		Find(nil).All(&list); err != nil {
		return nil
	}
	return list
}