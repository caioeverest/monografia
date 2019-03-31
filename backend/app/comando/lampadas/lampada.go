package lampadas

type Lampada struct {
	Ativo bool	`json:"ativo"`
}

func MontaComandoParaLampada(ativo bool) *Lampada {
	return &Lampada{ativo }
}