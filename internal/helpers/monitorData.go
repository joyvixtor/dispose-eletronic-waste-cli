package helpers

import (
	models "dispose-eletronic-waste-soft-version/internal/models"
	"fmt"
)

func ObterDadosMonitor() models.Monitor {
	fmt.Println("\n--- Preencha os dados para Desktop ou Notebook ---")
	return models.Monitor{
		Tipo:                  SelecionarOpcao("Tipo: ", models.TiposMonitor),
		TamanhoTela:           LerEntrada("Tamanho da Tela: "),
		Marca:                 LerEntrada("Marca: "),
		DescricaoComplementar: LerEntrada("Descrição Complementar: "),
		Classificacao:         SelecionarOpcao("Classificação: ", models.Classificacoes),
		ComponentesConstam:    SelecionarOpcao("Os componentes listados ainda constam no item?: ", models.ComponentesConstam),
		EstadoItem:            SelecionarOpcao("Estado do Item: ", models.Observacoes),
		SetorOrigem:           LerEntrada("Setor de Origem: "),
	}
}
