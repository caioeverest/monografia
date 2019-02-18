package repositorioDeSalas

type SalaEntity struct {
	Id       			string 		`bson:"_id,omitempty" json:"_id"`
	Campus				string		`bson:"campus,omitempty" json:"campus"`
	Bloco				string		`bson:"bloco,omitempty" json:"bloco"`
	Numero				string		`bson:"numero,omitempty" json:"numero"`
	Identificador		string		`bson:"identificador,omitempty" json:"identificador"`
}

func (d *SalaEntity) SetId(id string) {
	d.Id = id
}

func (d *SalaEntity) GetId() string {
	return d.Id
}

