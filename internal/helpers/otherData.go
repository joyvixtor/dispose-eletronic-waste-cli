package helpers

import (
	models "dispose-eletronic-waste-soft-version/internal/models"
	"fmt"
)

func ObterDadosOutros() models.Others {
	fmt.Println("\n--- Preencha os dados para Outros Itens de Informática ---")
	return models.Others{
		TipoEquipamento:       SelecionarOpcao("Tipo de Equipamento: ", models.TiposEquipamentoOutros),
		Marca:                 SelecionarOpcao("Marca: ", models.MarcasOutros),
		DescricaoComplementar: LerEntrada("Descrição Complementar: "),
		Classificacao:         SelecionarOpcao("Classificação: ", models.Classificacoes),
		ComponentesConstam:    SelecionarOpcao("Os componentes listados ainda constam no item?: ", models.ComponentesConstam),
		EstadoItem:            SelecionarOpcao("Estado do Item: ", models.Observacoes),
		SetorOrigem:           LerEntrada("Setor de Origem: "),
	}
}
