package helpers

import (
	models "dispose-eletronic-waste-soft-version/internal/models"
	"fmt"
)

func ObterDadosImpressora() models.Printer {
	fmt.Println("\n--- Preencha os dados para Impressora ---")
	return models.Printer{
		TipoImpressora:        SelecionarOpcao("Tipo: ", models.TiposImpressora),
		Marca:                 LerEntrada("Marca: "),
		DescricaoComplementar: LerEntrada("Descrição Complementar: "),
		Classificacao:         SelecionarOpcao("Classificação: ", models.Classificacoes),
		ComponentesConstam:    SelecionarOpcao("Os componentes listados ainda constam no item?: ", models.ComponentesConstam),
		EstadoItem:            SelecionarOpcao("Estado do Item: ", models.Observacoes),
		SetorOrigem:           LerEntrada("Setor de Origem: "),
	}
}
