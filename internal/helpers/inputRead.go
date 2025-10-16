package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LerEntrada(input string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(input)
	text, _ := reader.ReadString('\n')

	return strings.TrimSpace(text)
}

func SelecionarOpcao(prompt string, opcoes []string) string {
	fmt.Printf("\n%s:\n", prompt)
	for i, opt := range opcoes {
		fmt.Printf("%d. %s\n", i+1, opt)
	}

	for {
		escolhaStr := LerEntrada("Escolha uma opção: ")
		escolha, err := strconv.Atoi(escolhaStr)
		if err == nil && escolha > 0 && escolha <= len(opcoes) {
			return opcoes[escolha-1]
		}
		fmt.Println("Opção inválida. Tente novamente.")
	}
}
