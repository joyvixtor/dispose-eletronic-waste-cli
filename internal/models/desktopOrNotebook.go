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
	"até 100 MHz",
	"de 101MHz até 200MHz",
	"de 201MHz até 300MHz",
	"de 301MHz até 400MHz",
	"de 401MHz até 500MHz",
	"de 501MHz até 600MHz",
	"de 601MHz até 700MHz",
	"de 701MHz até 800MHz",
	"de 801MHz até 900MHz",
	"de 901MHz até 1GHz",
}

var Marcas = []string{
	"3CAM", "Abit", "Acer", "Advent", "AMD", "Apple", "ASI", "ASUS", "AT&T",
	"Compaq", "Creative", "Dell", "Digicom", "ECS", "Epson", "Everex",
	"Gateway", "Goldstar", "HP", "IBM", "Intel", "Itautec", "Lenovo",
	"LG", "Lite-On", "Microsoft", "MSI", "NEC", "Novell", "Nvidia",
	"Panasonic", "PCChips", "Philips", "Phitronics", "Positivo", "Samsung",
	"Sony", "Toshiba", "VIA", "ViewSonic", "Xerox", "Zebra", "Outra",
}
