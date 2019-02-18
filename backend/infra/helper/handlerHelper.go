package helper

type responseBuilder struct {
	Mensagem 	interface{}	`json:"mensagem"`
	Codigo		int		`json:"codigo"`
}

func ResponseBuilder(mensagem interface{}, codigo int) *responseBuilder{
	return &responseBuilder{
		Mensagem: mensagem,
		Codigo: codigo,
	}
}
