package helper

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

func Filler(s string) bson.ObjectId {
	key := s
	final := fmt.Sprintf("%12v", key)
	return bson.ObjectId(final)
}

func Binder(out interface{}, in map[string]interface{}) error {
	for k, v := range in {
		err := Escreve(out, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func Escreve(obj interface{}, name string, value interface{}) error {
	valorDaEstrutura := reflect.ValueOf(obj).Elem()
	valorDoCampo := valorDaEstrutura.FieldByName(name)

	if !valorDoCampo.IsValid() {
		return nil//fmt.Errorf("Campo [%s] não existe nesta estrutura", name)
	}

	if !valorDoCampo.CanSet() {
		return fmt.Errorf("Não consegue ver valor no campo [%s]", name)
	}

	tipoDoCampo := valorDoCampo.Type()
	val := reflect.ValueOf(value)
	if tipoDoCampo != val.Type() {
		return errors.New("O tipo que foi definido não bate com o da estrutura")
	}

	valorDoCampo.Set(val)
	return nil
}