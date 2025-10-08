// criar um arquivo utils/inputs.go pra colocar as funções de ler entrada do usuário e evitar repetição de código em cada arquivo
package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewReader(os.Stdin)

// ler a opção do menu e converte pra int pra usar no switch case
func LerOpcao() int {
	input, _ := scanner.ReadString('\n')
	input = strings.TrimSpace(input)
	opcao, err := strconv.Atoi(input)
	if err != nil {
		return -1 //retorna -1 se a conversão falhar, assim o switch case pode tratar como opção inválida e cair no default
	}
	return opcao
}

// ler uma string do terminal e remover espaços em branco pra já ajudar na limpeza do input
func LerString() string {
	input, _ := scanner.ReadString('\n')
	return strings.TrimSpace(input)
}
