package models

type Monitor struct {
	Tombamento            string
	Tipo                  string
	TamanhoTela           string
	Marca                 string
	DescricaoComplementar string
	Classificacao         string
	ComponentesConstam    string
	EstadoItem            string
	SetorOrigem           string
}

var TiposMonitor = []string{
	"CRT",
	"LCD",
	"OLED",
	"MICROLED",
	"LED",
	"Plasma",
	"Outro",
}
