package controllers

import (
	"fmt"
	"strings"
	"trabalho/database"
	"trabalho/models"
	"trabalho/utils"
	"trabalho/views"
)

// a função CadastrarCliente vai gerenciar o fluxo de cadastro de um cliente
// ela orquestra a interação entre a View e o Model
func CadastrarCliente() {
	//1. usa a função de dentro da viiew pra solicitar os dados do cliente
	nome, login, tipo, telefone, email, identificacao, endereco := views.SolicitarDadosCliente()

	//tava dando erro nas formatações pq as strings vinham com o '\n' do enter, tem que "limpar" elas
	nome = strings.TrimSpace(nome)
	login = strings.TrimSpace(login)
	tipo = strings.TrimSpace(tipo)
	telefone = strings.TrimSpace(telefone)
	email = strings.TrimSpace(email)
	identificacao = strings.TrimSpace(identificacao)
	endereco = strings.TrimSpace(endereco)

	//2. depois de pegar os dados, manda pro model criar um novo cliente usando o construtor recebendo os parâmetros
	novoCliente := models.NewCliente(nome, login, tipo, telefone, email, identificacao, endereco)

	//3. salva o cliente no banco
	database.SalvarCliente(novoCliente)

	//4. printa a confirmação de que cadastrou o cliente
	views.ExibirCliente(novoCliente)
}

// usa a função de listar clientes da view
func ListarClientes() {
	clientes := database.ClientesDB     //pega os dados do banco
	views.ExibirListaClientes(clientes) //manda pra view exibir a lista
}

// função pra editar os dados do cliente
func EditarCliente() {
	// 1. pede a identificação que é a chave pra achar o cliente no banco
	identificacao := views.SolicitarIdentificacao("editar")

	// 2. busca o cliente no banco verificando se existe
	cliente, existe := database.ClientesDB[identificacao]
	if !existe {
		fmt.Print("Cliente não encontrado.")
		return
	}

	// 3. exibe o cliente que achou e solicita os novos dados
	views.ExibirCliente(cliente)
	nome, login, telefone, email, endereco := views.SolicitarNovosDadosCliente(cliente)

	// 4. atualiza os dados (lembrando que se o usuário apertar enter, mantém o valor que já tava)
	if nome != "" {
		cliente.SetNome(nome)
	}
	if login != "" {
		cliente.SetLogin(login)
	}
	if telefone != "" {
		cliente.SetTelefone(telefone)
	}
	if email != "" {
		cliente.SetEmail(email)
	}
	if endereco != "" {
		cliente.SetEndereco(endereco)
	}

	fmt.Print("Cliente atualizado com sucesso!")
	views.ExibirCliente(cliente)
}

// função pra deletar um cliente
func DeletarCliente() {
	// 1. pede a identificação que é a chave pra achar o cliente no banco
	identificacao := views.SolicitarIdentificacao("deletar")

	// 2. busca o cliente no banco verificando se existe
	cliente, existe := database.ClientesDB[identificacao]
	if !existe {
		fmt.Print("Cliente não encontrado.")
		return
	}

	// 3. confirma se realmente quer deletar chamando a função da view
	if views.ConfirmarExclusaoCliente(cliente) {
		// 4. se voltou como true, deleta o cliente do banco
		delete(database.ClientesDB, identificacao)
		fmt.Print("Cliente excluído com sucesso.")
	} else {
		fmt.Print("Operação de exclusão cancelada.") //se não confirmou, cancela a exclusão
	}
}

func GerenciarClientes() {
	for {
		views.MostrarSubMenu("Clientes")
		opcao := utils.LerOpcao()

		switch opcao {
		case 1:
			CadastrarCliente()
		case 2:
			ListarClientes()
		case 3:
			EditarCliente()
		case 4:
			DeletarCliente()
		case 0:
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
