package builders

import (
	models "dispose-eletronic-waste-soft-version/internal/models"
	"fmt"
	"log"
	"time"

	"github.com/xuri/excelize/v2"
)

func GerarArquivos(
	desktops []models.DesktopOrNotebook,
	notebooks []models.DesktopOrNotebook,
	monitors []models.Monitor,
) {
	if len(desktops) == 0 && len(notebooks) == 0 {
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

	f.DeleteSheet("Sheet1")
	if err := f.SaveAs(nomeXLSX); err != nil {
		log.Fatal("Falha ao salvar o arquivo XLSX:", err)
	}
	fmt.Printf("\nPlanilha '%s' gerada com sucesso!\n", nomeXLSX)
}
