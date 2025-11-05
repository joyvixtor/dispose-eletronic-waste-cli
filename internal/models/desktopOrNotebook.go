package models

type DesktopOrNotebook struct {
	Tombamento            string
	Processador           string
	Clock                 string
	RAM                   string
	Armazenamento         string
	Marca                 string
	DescricaoComplementar string
	Classificacao         string
	ComponentesConstam    string
	EstadoItem            string
	SetorOrigem           string
}

var VelocidadesProcessador = []string{
	"até 1.0 GHz",
	"de 1.0 GHz até 1.5 GHz",
	"de 1.5 GHz até 2.0 GHz",
	"de 2.0 GHz até 2.5 GHz",
	"de 2.5 GHz até 3.0 GHz",
	"de 3.0 GHz até 3.5 GHz",
	"de 3.5 GHz até 4.0 GHz",
	"de 4.0 GHz até 4.5 GHz",
	"de 4.5 GHz até 5.0 GHz",
	"de 5.0 GHz até 5.5 GHz",
	"de 5.5 GHz até 6.0 GHz",
	"Maior que 6.0 GHz",
}

var Marcas = []string{
	"3CAM", "Abit", "Acer", "Advent", "AMD", "Apple", "ASI", "ASUS", "AT&T",
	"Compaq", "Creative", "Dell", "Digicom", "ECS", "Epson", "Everex",
	"Gateway", "Goldstar", "HP", "IBM", "Intel", "Itautec", "Lenovo",
	"LG", "Lite-On", "Microsoft", "MSI", "NEC", "Novell", "Nvidia",
	"Panasonic", "PCChips", "Philips", "Phitronics", "Positivo", "Samsung",
	"Sony", "Toshiba", "VIA", "ViewSonic", "Xerox", "Zebra", "Outra",
}
