package helpers

import (
	DesktopOrNotebook "dispose-eletronic-waste-soft-version/internal/models"
	models "dispose-eletronic-waste-soft-version/internal/models"
	"fmt"
)

func ObterDadosDesktopOuNotebook() DesktopOrNotebook.DesktopOrNotebook {
	fmt.Println("\n--- Preencha os dados para Desktop ou Notebook ---")
	return DesktopOrNotebook.DesktopOrNotebook{
		Processador:           LerEntrada("Processador: "),
		Clock:                 SelecionarOpcao("Velocidade: ", models.VelocidadesProcessador),
		RAM:                   LerEntrada("RAM: "),
		Armazenamento:         LerEntrada("Armazenamento: "),
		Marca:                 SelecionarOpcao("Marca: ", models.Marcas),
		DescricaoComplementar: LerEntrada("Descrição Complementar: "),
		Classificacao:         SelecionarOpcao("Classificação: ", models.Classificacoes),
		ComponentesConstam:    SelecionarOpcao("Os componentes listados ainda constam no item?: ", models.ComponentesConstam),
		EstadoItem:            SelecionarOpcao("Estado do Item: ", models.Observacoes),
		SetorOrigem:           LerEntrada("Setor de Origem: "),
	}
}
