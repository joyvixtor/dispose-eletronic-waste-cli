package main

import (
	builders "dispose-eletronic-waste-soft-version/internal/builders"
	helpers "dispose-eletronic-waste-soft-version/internal/helpers"
	DesktopOrNotebook "dispose-eletronic-waste-soft-version/internal/models"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var desktops []DesktopOrNotebook.DesktopOrNotebook
	var notebooks []DesktopOrNotebook.DesktopOrNotebook

	for {
		fmt.Println("\n--- Gerador de Planilha para Descarte de Bens ---")
		fmt.Println("Escolha o tipo de item que deseja adicionar:")
		fmt.Println("1. Desktop/Notebook")
		fmt.Println("2. Monitor")
		fmt.Println("3. Impressora")
		fmt.Println("4. Outros Itens de Informática")

		escolha := helpers.LerEntrada("Digite o número da opção desejada (ou 'sair' para encerrar): ")

		var tipoItem string
		var dadosBase interface{}

		switch escolha {
		case "1":
			tipoItem = helpers.LerEntrada("Digite 'D' para Desktop ou 'N' para Notebook: ")
			tipoItem = strings.ToUpper(tipoItem[:1]) + strings.ToLower(tipoItem[1:])
			if tipoItem != "D" && tipoItem != "N" {
				fmt.Println("Opção inválida. Tente novamente.")
				continue
			}
			dadosBase = helpers.ObterDadosDesktopOuNotebook()
		default:
			fmt.Println("Opção inválida. Por favor, tente novamente")
			continue
		}

		qtdeStr := helpers.LerEntrada("Quantos itens deste tipo você deseja adicionar? ")
		quantidade, err := strconv.Atoi(qtdeStr)
		if err != nil {
			fmt.Println("Quantidade inválida. Por favor, insira um número.")
			continue
		}

		for i := 0; i < quantidade; i++ {
			tombamento := helpers.LerEntrada(fmt.Sprintf("Digite o tombamento do item %d: ", i+1))

			switch tipoItem {
			case "D":
				item := dadosBase.(DesktopOrNotebook.DesktopOrNotebook)
				item.Tombamento = tombamento
				desktops = append(desktops, item)
			case "N":
				item := dadosBase.(DesktopOrNotebook.DesktopOrNotebook)
				item.Tombamento = tombamento
				notebooks = append(notebooks, item)
			}
		}

		adicionarMais := helpers.LerEntrada("Deseja adicionar mais itens? (s/n): ")
		if strings.ToLower(adicionarMais) != "s" {
			break
		}
	}

	builders.GerarArquivos(desktops, notebooks)
	fmt.Println("\nPrograma encerrado.")
}
