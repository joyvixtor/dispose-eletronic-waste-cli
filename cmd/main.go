package main

import (
	builders "dispose-eletronic-waste-soft-version/internal/builders"
	helpers "dispose-eletronic-waste-soft-version/internal/helpers"
	models "dispose-eletronic-waste-soft-version/internal/models"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var desktops []models.DesktopOrNotebook
	var notebooks []models.DesktopOrNotebook
	var monitors []models.Monitor
	var printers []models.Printer
	var others []models.Others
	var servidor models.Servidor

	for {
		fmt.Println("\n--- Gerador de Planilha para Descarte de Bens ---")
		fmt.Println("Escolha o tipo de item que deseja adicionar:")
		fmt.Println("1. Desktop/Notebook")
		fmt.Println("2. Monitor")
		fmt.Println("3. Impressora")
		fmt.Println("4. Outros Itens de Informática")

		escolha := helpers.LerEntrada("Digite o número da opção desejada: ")

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
		case "2":
			tipoItem = "Monitor"
			dadosBase = helpers.ObterDadosMonitor()
		case "3":
			tipoItem = "Impressora"
			dadosBase = helpers.ObterDadosImpressora()
		case "4":
			tipoItem = "Outros"
			dadosBase = helpers.ObterDadosOutros()
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
				item := dadosBase.(models.DesktopOrNotebook)
				item.Tombamento = tombamento
				desktops = append(desktops, item)
			case "N":
				item := dadosBase.(models.DesktopOrNotebook)
				item.Tombamento = tombamento
				notebooks = append(notebooks, item)
			case "Monitor":
				item := dadosBase.(models.Monitor)
				item.Tombamento = tombamento
				monitors = append(monitors, item)
			case "Impressora":
				item := dadosBase.(models.Printer)
				item.Tombamento = tombamento
				printers = append(printers, item)
			case "Outros":
				item := dadosBase.(models.Others)
				item.Tombamento = tombamento
				others = append(others, item)
			}
		}

		adicionarMais := helpers.LerEntrada("Deseja adicionar mais itens? (s/n): ")
		if strings.ToLower(adicionarMais) != "s" {

			fmt.Println("Escolha o servidor responsável por este documento:")
			fmt.Println("1. Christian")
			fmt.Println("2. Renan")
			fmt.Println("3. Emanuel")

			escolha = helpers.LerEntrada("Digite o número da opção desejada: ")

			switch escolha {
			case "1":
				servidor = models.ServidorChristian
			case "2":
				servidor = models.ServidorRenan
			case "3":
				servidor = models.ServidorEmanuel
			default:
				fmt.Println("Opção inválida. Por favor, tente novamente.")
				continue
			}
			break
		}
	}

	builders.GerarArquivos(desktops, notebooks, monitors, printers, servidor, others)
	fmt.Println("\nPrograma encerrado.")
}
