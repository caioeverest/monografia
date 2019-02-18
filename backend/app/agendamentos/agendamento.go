package agendamentos

import (
	"errors"
	"github.com/caioever/monografia/backend/adapters/repositorioDeAgendamentos"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	EM_ESPERA	=	"EM_ESPERA"
	INICIADO	= 	"INICIADO"
	FINALIZADO	=	"FINALIZADO"
)

var repository = repositorioDeAgendamentos.Init()

type Agendamento struct {
	Id				string			`json:"id,omitempty"`
	Sala 			string			`json:"sala"`
	Professor 		string			`json:"professor"`
	Status			string			`json:"status"`
	Inicio 			time.Time		`json:"inicio"`
	Fim 			time.Time		`json:"fim"`
}

type AgendamentoFrendly struct {
	Id				string			`json:"id,omitempty"`
	Sala 			string			`json:"sala"`
	Professor 		string			`json:"professor"`
	Inicio 			string			`json:"inicio"`
	Fim 			string			`json:"fim"`
	Status			string			`json:"status"`
}

func Builder(id, sala, professor, status string, inicio, fim time.Time) *Agendamento {
	return &Agendamento{
		Id:			id,
		Sala: 		sala,
		Professor: 	professor,
		Status:		status,
		Inicio: 	inicio,
		Fim: 		fim,
	}
}

func EntitySimplifier(dbEntity repositorioDeAgendamentos.AgendamentoEntity) *Agendamento{
	return Builder(dbEntity.Id.Hex(), dbEntity.Sala, dbEntity.Professor,
		dbEntity.Status, dbEntity.Inicio, dbEntity.Fim)
}

func TrazerTodosAgendamentos() (interface{}, error) {
	var output []AgendamentoFrendly
	dadosBd := repository.RecuperaAtivosEFormataAgendamentos()
	for _, agendamento := range dadosBd {
		this := AgendamentoFrendly{
			Id:			agendamento.Id.Hex(),
			Sala: 		agendamento.Sala,
			Professor: 	agendamento.Professor,
			Inicio: 	agendamento.Inicio.Format("_2/Jan 15:04"),
			Fim: 		agendamento.Fim.Format("_2/Jan 15:04"),
			Status:		agendamento.Status,
		}
		output = append(output, this)
	}
	return output, nil
}

func TrazerAgendamentosEntre(inicio, fim time.Time) []Agendamento {
	agendamentos := repository.RetornarAgendamentoEntre(inicio, fim)
	var listaDeAgendamentosFormatados []Agendamento
	for _, agendamento := range agendamentos {
		agendamentoFormatado := Builder(agendamento.Id.Hex(), agendamento.Sala, agendamento.Professor, agendamento.Status, agendamento.Inicio, agendamento.Fim)
		listaDeAgendamentosFormatados = append(listaDeAgendamentosFormatados, *agendamentoFormatado)
	}
	return listaDeAgendamentosFormatados
}

func CriarNovoAgendamento(n *Agendamento) error {
	if n.Inicio.After(n.Fim) {
		return errors.New("fim antes od inicio")
	}
	agendamento := &repositorioDeAgendamentos.AgendamentoEntity{
		Sala:     	n.Sala,
		Professor:	n.Professor,
		Status:		EM_ESPERA,
		Inicio:   	n.Inicio,
		Fim:      	n.Fim,
	}
	if err := repository.Save(agendamento); err != nil {
		return err
	}
	return nil
}

func RecuperaAgendamentoPeloId(id string) (interface{}, error) {
	agendamentoEntity, err := repository.FindById(bson.ObjectIdHex(id))
	if err != nil {
		return nil, err
	}

	return agendamentoEntity, nil
}

func RecuperaAgendamentoPeloIdSimplificado(id string) (*Agendamento, error) {
	dataRaw, err := repository.FindById(bson.ObjectIdHex(id))
	if err != nil {
		return nil, err
	}
	agendamentoEntity := &repositorioDeAgendamentos.AgendamentoEntity{}
	dataHandler := dataRaw.(bson.M)
	bsonBytes, _ := bson.Marshal(dataHandler)
	if err := bson.Unmarshal(bsonBytes, &agendamentoEntity); err != nil {
		return nil, err
	}
	agendamento := Builder(agendamentoEntity.Id.Hex(),
		agendamentoEntity.Sala,
		agendamentoEntity.Professor,
		agendamentoEntity.Status,
		agendamentoEntity.Inicio,
		agendamentoEntity.Fim)

	return agendamento, nil
}

func AtualizaAgendamento(n echo.Map, id string) error {
	return repository.Update(n, bson.ObjectIdHex(id))
}

func DeletaAgendamento(id string) error {
	return repository.Delete(bson.ObjectIdHex(id))
}

func Monitorar(query map[string]interface{}) []repositorioDeAgendamentos.AgendamentoEntity {
	return repository.Monitorar(query)
}

func AtualizarStatusDoAgendamento(id string, novoStatus string) error {
	change := make(map[string]interface{})
	change["status"] = novoStatus
	return repository.Update(change, bson.ObjectIdHex(id))
}