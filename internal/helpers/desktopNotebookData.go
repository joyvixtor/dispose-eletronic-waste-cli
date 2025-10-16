package helpers

import (
	DesktopOrNotebook "dispose-eletronic-waste-soft-version/internal/models"
	"fmt"
)

func ObterDadosDesktopOuNotebook() DesktopOrNotebook.DesktopOrNotebook {
	fmt.Println("\n--- Preencha os dados para Desktop ou Notebook ---")
	return DesktopOrNotebook.DesktopOrNotebook{
		Processador:           LerEntrada("Processador: "),
		Clock:                 LerEntrada("Clock: "),
		RAM:                   LerEntrada("RAM: "),
		Armazenamento:         LerEntrada("Armazenamento: "),
		Marca:                 LerEntrada("Marca: "),
		DescricaoComplementar: LerEntrada("Descrição Complementar: "),
		Classificacao:         LerEntrada("Classificação: "),
		ComponentesConstam:    LerEntrada("Componentes que Constam: "),
		EstadoItem:            LerEntrada("Estado do Item: "),
		SetorOrigem:           LerEntrada("Setor de Origem: "),
	}
}
