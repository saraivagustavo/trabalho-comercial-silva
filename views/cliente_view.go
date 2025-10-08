package views

import (
	"bufio"           //para ler strings
	"fmt"             //pra imprimir na tela
	"os"              //ler do terminal
	"strings"         //manipular strings
	"trabalho/models" //pegar as structs
	"trabalho/utils"  //usar a função utilitária de ler string
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

// SolicitarIdentificacao pede o cpf ou cnpj para encontrar um cliente específico que vai ser editado ou deletado
func SolicitarIdentificacao(acao string) string {
	fmt.Printf("\n--- %s Cliente ---\n", acao) //ação vai ser "Editar" ou "Deletar", vai puxar a string que foi passada como parâmetro
	fmt.Print("Digite o CPF/CNPJ do cliente: ")
	return utils.LerString()
}

// SolicitarNovosDadosCliente pra atualizar os dados do cliente, se não quiser mudar todos os dados, se o usuário apertar enter, mantém o valor que já tava
func SolicitarNovosDadosCliente(cliente *models.Cliente) (nome, login, telefone, email, endereco string) {
	fmt.Println("\nDigite os novos dados. Pressione ENTER para manter o valor atual.")

	// nome
	fmt.Printf("Nome atual: %s\nNovo nome: ", cliente.GetNome())
	nome = utils.LerString()

	// login
	fmt.Printf("Login atual: %s\nNovo login: ", cliente.GetLogin())
	login = utils.LerString()

	// telefone
	fmt.Printf("Telefone atual: %s\nNovo telefone: ", cliente.GetTelefone())
	telefone = utils.LerString()

	// email
	fmt.Printf("Email atual: %s\nNovo email: ", cliente.GetEmail())
	email = utils.LerString()

	// endereço
	fmt.Printf("Endereço atual: %s\nNovo endereço: ", cliente.GetEndereco())
	endereco = utils.LerString()

	return
}

// ConfirmarExclusao só pra segurança de não deletar o cliente errado sem querer
func ConfirmarExclusaoCliente(cliente *models.Cliente) bool {
	fmt.Printf("Tem certeza que deseja excluir o cliente: %s (CPF/CNPJ: %s)?\n", cliente.GetNome(), cliente.GetIdentificacao())
	fmt.Print("Digite 's' para confirmar: ")
	resposta := utils.LerString()
	return strings.ToLower(resposta) == "s" //se for 's' isso da true e confirma a exclusão
}
