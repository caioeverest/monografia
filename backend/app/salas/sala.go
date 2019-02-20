package salas

import (
	"github.com/caioever/monografia/backend/adapters/repositorioDeSalas"
	"github.com/caioever/monografia/backend/app/agendamentos"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Sala struct {
	Campus        string `json:"campus"`
	Bloco         string `json:"bloco"`
	Numero        string `json:"numero"`
	Identificador string `json:"identificador"`
}

var repository = repositorioDeSalas.Init()

func TrazerTodasAsSalas() []Sala {
	var output []Sala
	dadosBanco := repository.TodasAsSalas()
	for _, sala := range dadosBanco {
		this := Sala{
			Campus:        sala.Campus,
			Bloco:         sala.Bloco,
			Numero:        sala.Numero,
			Identificador: sala.Identificador,
		}
		output = append(output, this)
	}
	return output
}

func Existe(id string) bool {
	if sala, _ := EncontrarSalaPorIdentificador(id); sala != nil {
		return true
	}
	return false
}

func TrazerIdentificadoresDeTodasAsSalas() ([]string, error) {
	identificadores, err := repository.RecuperarIdentificadores()
	if err != nil {
		return nil, err
	}
	return identificadores, nil
}

func EncontrarSalaPorIdentificador(identificador string) (*Sala, error) {
	r, err := repository.EncontraPeloIdentificador(identificador)
	if err != nil {
		return &Sala{}, err
	}
	out := Sala{
		Campus:        r.Campus,
		Bloco:         r.Bloco,
		Numero:        r.Numero,
		Identificador: r.Identificador,
	}
	return &out, nil
}

func CriarNovaSala(n *Sala) error {
	sala := &repositorioDeSalas.SalaEntity{
		Campus:        n.Campus,
		Bloco:         n.Bloco,
		Numero:        n.Numero,
		Identificador: n.Identificador,
	}
	id := n.Identificador
	sala.SetId(bson.ObjectId(id).Hex())

	return repository.Save(sala)
}

func RecuperaSalaPeloId(id string) (interface{}, error) {
	output, err := repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func AtualizaSala(n echo.Map, id string) error {
	return repository.Update(n, id)
}

func SalasDisponiveis(inicio, fim time.Time) []string {
	var salasDisponiveis []string
	salasOcupadas := SalasOcupadas(inicio, fim)
	todasAsSalas := TrazerTodasAsSalas()

	for _, sala := range todasAsSalas {
		if _, ok := salasOcupadas[sala.Identificador]; !ok {
			salasDisponiveis = append(salasDisponiveis, sala.Identificador)
		}
	}
	return salasDisponiveis
}

func SalasOcupadas(inicio, fim time.Time) map[string]string {
	mapa := make(map[string]string)
	agem := agendamentos.TrazerAgendamentosEntre(inicio, fim)

	for _, agendamento := range agem {
		mapa[agendamento.Sala] = agendamento.Id
	}
	return mapa
}

func DeletaSala(id string) error {
	return repository.Delete(id)
}
