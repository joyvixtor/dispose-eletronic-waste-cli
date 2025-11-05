package models

import "time"

type Servidor struct {
	ResponsavelTecnico string
	Siape              string
	DataAvaliacao      string
}

var ServidorRenan = Servidor{
	ResponsavelTecnico: "Renan Mousinho Aquino",
	Siape:              "2125855",
	DataAvaliacao:      time.Now().Format("02/01/2006"),
}

var ServidorChristian = Servidor{
	ResponsavelTecnico: "Christian Douglas Ramos dos Santos",
	Siape:              "1426068",
	DataAvaliacao:      time.Now().Format("02/01/2006"),
}

var ServidorEmanuel = Servidor{
	ResponsavelTecnico: "Emanuel Arruda da Silva",
	Siape:              "3124849",
	DataAvaliacao:      time.Now().Format("02/01/2006"),
}
