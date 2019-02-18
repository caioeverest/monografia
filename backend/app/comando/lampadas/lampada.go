package lampadas

type Lampada struct {
	Ativo bool
}

func MontaComandoParaLampada(ativo bool) *Lampada {
	return &Lampada{ativo }
}