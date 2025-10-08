package views

import (
	"bufio"           //para ler strings
	"fmt"             //pra imprimir na tela
	"os"              //ler do terminal
	"trabalho/models" //pegar as structs
)

// a função lê os dados do cliente a partir do terminal e retorna os valores
func SolicitarDadosCliente() (nome, login, tipo, telefone, email, identificacao, endereco string) {
	//funciona similar ao scanner do java
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- CADASTRO DE NOVO CLIENTE ---")
	fmt.Print("Nome Completo: ")
	nome, _ = scanner.ReadString('\n') //a ReadString lê até encontrar o caractere '\n' de enter, retorna 2 valores, a string e um erro, o "_underline" ignora o erro

	fmt.Print("Login: ")
	login, _ = scanner.ReadString('\n')

	tipo = "cliente"

	fmt.Print("Telefone DDD + número (apenas números): ")
	telefone, _ = scanner.ReadString('\n')

	fmt.Print("Email: ")
	email, _ = scanner.ReadString('\n')

	fmt.Print("CPF/CNPJ (apenas números): ")
	identificacao, _ = scanner.ReadString('\n')

	fmt.Print("Endereço: ")
	endereco, _ = scanner.ReadString('\n')
	fmt.Println("-------------------------------")

	return
}

// a função exibe os dados de um cliente passado como parâmetro
func ExibirCliente(cliente *models.Cliente) {
	fmt.Println("\n--- CLIENTE CADASTRADO ---")
	fmt.Println(cliente.ToString())
	fmt.Println("--------------------------")
}

// função pra exibir todos os clientes cadastrados
func ExibirListaClientes(clientes map[string]*models.Cliente) { //usar map aqui porque a database armazena os clientes em um map
	fmt.Println("\n--- LISTA DE CLIENTES ---")
	if len(clientes) == 0 {
		fmt.Println("Nenhum cliente cadastrado.")
		return
	}
	for _, cliente := range clientes { //loop pra percorrer o map e exibir cada cliente
		fmt.Println(cliente.ToString())
		fmt.Println("-------------------------")
	}
}
