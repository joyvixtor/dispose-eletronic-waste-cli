package models

type Printer struct {
	Tombamento            string
	TipoImpressora        string
	Marca                 string
	DescricaoComplementar string
	Classificacao         string
	ComponentesConstam    string
	EstadoItem            string
	SetorOrigem           string
}

var TiposImpressora = []string{
	"Jato de Tinta",
	"Laser",
	"Matricial",
	"Multifuncional",
	"3D",
	"Outra",
}

var MarcasImpressora = []string{
	"3CAM",
	"Abit",
	"Acer",
	"Advent",
	"AMD",
	"Apple",
	"ASI",
	"ASUS",
	"AT&T",
	"Compaq",
	"Creative",
	"Dell",
	"Digicom",
	"ECS",
	"Epson",
	"Everex",
	"Gateway",
	"Goldstar",
	"HP",
	"IBM",
	"Intel",
	"Itautec",
	"Lenovo",
	"LG",
	"Lite-On",
	"Microsoft",
	"MSI",
	"NEC",
	"Novell",
	"Nvidia",
	"Panasonic",
	"PCChips",
	"Philips",
	"Phitronics",
	"Positivo",
	"Samsung",
	"Sony",
	"Toshiba",
	"Trident Microsystems",
	"Tsunami",
	"TVSE Electronics",
	"Twinhead",
	"UMC",
	"Veo Intl",
	"VIA",
	"ViewSonic",
	"Xerox",
	"Zebra",
	"OUTROS",
}
