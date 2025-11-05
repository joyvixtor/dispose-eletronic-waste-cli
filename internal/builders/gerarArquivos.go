package builders

import (
	models "dispose-eletronic-waste-soft-version/internal/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/xuri/excelize/v2"
)

func ajustarLarguraColunas(f *excelize.File, sheetName string, headers []string) {
	for i, header := range headers {
		col, _ := excelize.ColumnNumberToName(i + 1)
		width := float64(len(header)) + 8
		f.SetColWidth(sheetName, col, col, width)
	}
}

func adicionarLinhaServidor(f *excelize.File, sheetName string, servidor models.Servidor, numColunas int) {
	textoServidor := fmt.Sprintf("Responsável técnico: %s  SIAPE: %s  Data de avaliação: %s",
		servidor.ResponsavelTecnico,
		servidor.Siape,
		servidor.DataAvaliacao)

	f.SetCellValue(sheetName, "A1", textoServidor)

	ultimaColuna, _ := excelize.ColumnNumberToName(numColunas)
	f.MergeCell(sheetName, "A1", fmt.Sprintf("%s1", ultimaColuna))

	style, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 11,
		},
	})
	f.SetCellStyle(sheetName, "A1", fmt.Sprintf("%s1", ultimaColuna), style)

	f.SetRowHeight(sheetName, 1, 20)
}

func aplicarEstiloHeader(f *excelize.File, sheetName string, numColunas int) {
	styleHeader, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 11,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	ultimaColuna, _ := excelize.ColumnNumberToName(numColunas)
	f.SetCellStyle(sheetName, "A2", fmt.Sprintf("%s2", ultimaColuna), styleHeader)
	f.SetRowHeight(sheetName, 2, 30)
}

func aplicarEstiloDados(f *excelize.File, sheetName string, linha int, numColunas int) {
	styleDados, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	ultimaColuna, _ := excelize.ColumnNumberToName(numColunas)
	f.SetCellStyle(sheetName, fmt.Sprintf("A%d", linha), fmt.Sprintf("%s%d", ultimaColuna, linha), styleDados)
	f.SetRowHeight(sheetName, linha, 20)
}

func GerarArquivos(
	desktops []models.DesktopOrNotebook,
	notebooks []models.DesktopOrNotebook,
	monitors []models.Monitor,
	printers []models.Printer,
	servidor models.Servidor,
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

		adicionarLinhaServidor(f, "Desktop", servidor, len(headers))

		f.SetSheetRow("Desktop", "A2", &headers)
		ajustarLarguraColunas(f, "Desktop", headers)
		aplicarEstiloHeader(f, "Desktop", len(headers))

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
			linhaAtual := i + 3
			f.SetSheetRow("Desktop", fmt.Sprintf("A%d", linhaAtual), &row)
			aplicarEstiloDados(f, "Desktop", linhaAtual, len(headers))
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

		adicionarLinhaServidor(f, "Notebooks", servidor, len(headers))

		f.SetSheetRow("Notebooks", "A2", &headers)
		ajustarLarguraColunas(f, "Notebooks", headers)
		aplicarEstiloHeader(f, "Notebooks", len(headers))

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
			linhaAtual := i + 3
			f.SetSheetRow("Notebooks", fmt.Sprintf("A%d", linhaAtual), &row)
			aplicarEstiloDados(f, "Notebooks", linhaAtual, len(headers))
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

		adicionarLinhaServidor(f, "Monitors", servidor, len(headers))

		f.SetSheetRow("Monitors", "A2", &headers)
		ajustarLarguraColunas(f, "Monitors", headers)
		aplicarEstiloHeader(f, "Monitors", len(headers))

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
			linhaAtual := i + 3
			f.SetSheetRow("Monitors", fmt.Sprintf("A%d", linhaAtual), &row)
			aplicarEstiloDados(f, "Monitors", linhaAtual, len(headers))
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

		adicionarLinhaServidor(f, "Printers", servidor, len(headers))

		f.SetSheetRow("Printers", "A2", &headers)
		ajustarLarguraColunas(f, "Printers", headers)
		aplicarEstiloHeader(f, "Printers", len(headers))

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
			linhaAtual := i + 3
			f.SetSheetRow("Printers", fmt.Sprintf("A%d", linhaAtual), &row)
			aplicarEstiloDados(f, "Printers", linhaAtual, len(headers))
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

		adicionarLinhaServidor(f, "Others", servidor, len(headers))

		f.SetSheetRow("Others", "A2", &headers)
		ajustarLarguraColunas(f, "Others", headers)
		aplicarEstiloHeader(f, "Others", len(headers))

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
			linhaAtual := i + 3
			f.SetSheetRow("Others", fmt.Sprintf("A%d", linhaAtual), &row)
			aplicarEstiloDados(f, "Others", linhaAtual, len(headers))
		}
	}

	f.DeleteSheet("Sheet1")

	execPath, err := os.Executable()
	if err != nil {
		log.Fatal("Erro ao obter caminho do executável:", err)
	}
	dirExecutavel := filepath.Dir(execPath)

	caminhoCompleto := filepath.Join(dirExecutavel, nomeXLSX)

	if err := f.SaveAs(caminhoCompleto); err != nil {
		log.Fatal("Falha ao salvar o arquivo XLSX:", err)
	}
	fmt.Printf("\nPlanilha '%s' gerada com sucesso em: %s\n", nomeXLSX, dirExecutavel)
}
