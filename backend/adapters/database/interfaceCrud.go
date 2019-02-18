package database

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)


type Repository struct {
	collection 	string
	model		interface{}
}

func Inject(c string, m interface{}) *Repository {
	return &Repository{
		collection: c,
		model: m,
	}
}

func (repository *Repository) Save(estrutura interface{}) error {
	now := time.Now()
	if tt, ok := estrutura.(DefTimeInfo); ok {
		tt.SetCreated(now)
		tt.SetModified(now)
	}
	if especial, ok := estrutura.(DocEspecial); ok {
		if repository.Exists(especial.GetId()) {
			return fmt.Errorf("Já Exists este documeto")
		}
		_, err := Conectar().C(repository.collection).UpsertId(especial.GetId() ,estrutura)
		return err
	}
	err := Conectar().C(repository.collection).Insert(&estrutura)
	return err
}

func (repository *Repository) FindAll() (interface{}, error) {
	list := make([]interface{}, 0)
	err := Conectar().C(repository.collection).Find(nil).All(&list)
	return list, err
}

func (repository *Repository) Find(query bson.M) []interface{} {
	list := make([]interface{}, 0)
	if err := Conectar().C(repository.collection).
		Find(query).All(list); err != nil {
			return nil
	}
	return list
}

func (repository *Repository) FindById(chave interface{}) (interface{}, error) {
	mod := &repository.model
	err := Conectar().C(repository.collection).
		FindId(chave).
		One(mod)
	return *mod, err
}

func (repository *Repository) Exists(chave interface{}) bool {
	mod := &BaseDoc{}
	err := Conectar().C(repository.collection).FindId(chave).One(mod)
	if mod.GetId() == "" || err != nil {
		return false
	}
	return true
}

func (repository *Repository) Update(mudar map[string]interface{}, chave interface{}) error {
	mudar["_modified"] = time.Now()
	if !repository.Exists(chave) {
		log.Print(chave)
		return fmt.Errorf("Não Exists documento")
	}
	err := Conectar().C(repository.collection).
		Update(bson.M{"_id": chave}, bson.M{"$set": mudar})
	return err
}

func (repository *Repository) Delete(chave interface{}) error {
	return Conectar().C(repository.collection).RemoveId(chave)
}

func (repository *Repository) Collection() *mgo.Collection {
	return Conectar().C(repository.collection)
}