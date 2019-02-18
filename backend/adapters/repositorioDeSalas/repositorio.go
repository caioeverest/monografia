package repositorioDeSalas

import (
	"fmt"
	"github.com/caioever/monografia/backend/adapters/database"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type RepositorioDeSalas struct {
	*database.Repository
}

func Init() *RepositorioDeSalas {
	r := database.Inject("salas", SalaEntity{})
	return &RepositorioDeSalas{r}
}

func (repository RepositorioDeSalas) EncontraPeloIdentificador(indentificador string) (*SalaEntity, error) {
	mod := &SalaEntity{}
	query := make(bson.M)
	query["identificador"] = indentificador
	err := repository.Collection().Find(query).One(mod)
	if err != nil {
		log.Printf("Algo deu errado %s", err)
		return nil, fmt.Errorf("não foi encontrado sala com esse identificador, erro: %s", err.Error())
	}
	return mod, err
}

func (repository RepositorioDeSalas) RecuperarIdentificadores() ([]string, error) {
	var mod []SalaEntity
	var list []string
	err := repository.Collection().Find(nil).All(&mod)
	if err != nil {
		log.Printf("Algo deu errado %s", err)
		return nil, fmt.Errorf("não foi encontrado sala com esse identificador, erro: %s", err.Error())
	}
	for _, sala := range mod {
		list = append(list, sala.Identificador)
	}
	return list, err
}

func (repository *RepositorioDeSalas) TodasAsSalas() ([]SalaEntity) {
	var list []SalaEntity
	if err := repository.Collection().
		Find(nil).All(&list); err != nil {
		return nil
	}
	return list
}