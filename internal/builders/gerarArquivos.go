package builders

import (
	models "dispose-eletronic-waste-soft-version/internal/models"
	"fmt"
	"log"
	"time"

	"github.com/xuri/excelize/v2"
)

func ajustarLarguraColunas(f *excelize.File, sheetName string, headers []string) {
	for i, header := range headers {
		col, _ := excelize.ColumnNumberToName(i + 1)
		width := float64(len(header)) + 2
		f.SetColWidth(sheetName, col, col, width)
	}
}

func GerarArquivos(
	desktops []models.DesktopOrNotebook,
	notebooks []models.DesktopOrNotebook,
	monitors []models.Monitor,
	printers []models.Printer,
	others []models.Others,
) {
	if len(desktops) == 0 && len(notebooks) == 0 && len(monitors) == 0 && len(printers) == 0 && len(others) == 0 {
		return
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	nomeArquivoBase := fmt.Sprintf("descarte_de_bens_%s", timestamp)
	nomeXLSX := nomeArquivoBase + ".xlsx"

	f := excelize.NewFile()

	if len(desktops) > 0 {
		f.NewSheet("Desktop")
		headers := []string{
			"Nº Tombamento",
			"Processador",
			"Velocidade/Clock",
			"Memória RAM",
			"Armazenamento",
			"Marca",
			"Descrição",
			"Classificação",
			"Componentes Constam?",
			"Observações sobre o estado do bem",
			"Setor Origem",
		}
		f.SetSheetRow("Desktop", "A1", &headers)
		ajustarLarguraColunas(f, "Desktop", headers)
		for i, item := range desktops {
			row := []interface{}{
				item.Tombamento,
				item.Processador,
				item.Clock,
				item.RAM,
				item.Armazenamento,
				item.Marca,
				item.DescricaoComplementar,
				item.Classificacao,
				item.ComponentesConstam,
				item.EstadoItem,
				item.SetorOrigem,
			}
			f.SetSheetRow("Desktop", fmt.Sprintf("A%d", i+2), &row)
		}
	}

	if len(notebooks) > 0 {
		f.NewSheet("Notebooks")
		headers := []string{
			"Nº Tombamento",
			"Processador",
			"Velocidade/Clock",
			"Memória RAM",
			"Armazenamento",
			"Marca",
			"Descrição",
			"Classificação",
			"Componentes Constam?",
			"Observações sobre o estado do bem",
			"Setor Origem",
		}
		f.SetSheetRow("Notebooks", "A1", &headers)
		ajustarLarguraColunas(f, "Notebooks", headers)
		for i, item := range notebooks {
			row := []interface{}{
				item.Tombamento,
				item.Processador,
				item.Clock,
				item.RAM,
				item.Armazenamento,
				item.Marca,
				item.DescricaoComplementar,
				item.Classificacao,
				item.ComponentesConstam,
				item.EstadoItem,
				item.SetorOrigem,
			}
			f.SetSheetRow("Notebooks", fmt.Sprintf("A%d", i+2), &row)
		}
	}

	if len(monitors) > 0 {
		f.NewSheet("Monitors")
		headers := []string{
			"Nº Tombamento",
			"Tipo do monitor",
			"Tamanho da Tela",
			"Marca",
			"Descrição",
			"Classificação",
			"Componentes Constam?",
			"Observações sobre o estado do bem",
			"Setor Origem",
		}
		f.SetSheetRow("Monitors", "A1", &headers)
		ajustarLarguraColunas(f, "Monitors", headers)
		for i, item := range monitors {
			row := []interface{}{
				item.Tombamento,
				item.Tipo,
				item.TamanhoTela,
				item.Marca,
				item.DescricaoComplementar,
				item.Classificacao,
				item.ComponentesConstam,
				item.EstadoItem,
				item.SetorOrigem,
			}
			f.SetSheetRow("Monitors", fmt.Sprintf("A%d", i+2), &row)
		}
	}

	if len(printers) > 0 {
		f.NewSheet("Printers")
		headers := []string{
			"Nº Tombamento",
			"Tipo da Impressora",
			"Marca",
			"Descrição",
			"Classificação",
			"Componentes Constam?",
			"Observações sobre o estado do bem",
			"Setor Origem",
		}
		f.SetSheetRow("Printers", "A1", &headers)
		ajustarLarguraColunas(f, "Printers", headers)
		for i, item := range printers {
			row := []interface{}{
				item.Tombamento,
				item.TipoImpressora,
				item.Marca,
				item.DescricaoComplementar,
				item.Classificacao,
				item.ComponentesConstam,
				item.EstadoItem,
				item.SetorOrigem,
			}
			f.SetSheetRow("Printers", fmt.Sprintf("A%d", i+2), &row)
		}
	}

	if len(others) > 0 {
		f.NewSheet("Others")
		headers := []string{
			"Nº Tombamento",
			"Tipo de Equipamento",
			"Marca",
			"Descrição",
			"Classificação",
			"Componentes Constam?",
			"Observações sobre o estado do bem",
			"Setor Origem",
		}
		f.SetSheetRow("Others", "A1", &headers)
		ajustarLarguraColunas(f, "Others", headers)
		for i, item := range others {
			row := []interface{}{
				item.Tombamento,
				item.TipoEquipamento,
				item.Marca,
				item.DescricaoComplementar,
				item.Classificacao,
				item.ComponentesConstam,
				item.EstadoItem,
				item.SetorOrigem,
			}
			f.SetSheetRow("Others", fmt.Sprintf("A%d", i+2), &row)
		}
	}

	f.DeleteSheet("Sheet1")
	if err := f.SaveAs(nomeXLSX); err != nil {
		log.Fatal("Falha ao salvar o arquivo XLSX:", err)
	}
	fmt.Printf("\nPlanilha '%s' gerada com sucesso!\n", nomeXLSX)
}
