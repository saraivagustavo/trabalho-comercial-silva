package controllers

import (
	"fmt"
	"strings"
	"trabalho/database"
	"trabalho/models"
	"trabalho/utils"
	"trabalho/views"
)

// mesma lógica do cliente_controller
func CadastrarFuncionario() {
	// 1. view coleta os dados
	nome, login, tipo, telefone, email, cargo, matricula := views.SolicitarDadosFuncionario()

	//tem que ver como vai ser a senha depois
	senhaTeste := "123456"

	//limpar as strings também
	nome = strings.TrimSpace(nome)
	login = strings.TrimSpace(login)
	tipo = strings.TrimSpace(tipo)
	telefone = strings.TrimSpace(telefone)
	email = strings.TrimSpace(email)
	cargo = strings.TrimSpace(cargo)
	matricula = strings.TrimSpace(matricula)

	// 2. model cria o objeto
	novoFuncionario := models.NewFuncionario(nome, login, senhaTeste, tipo, telefone, email, cargo, matricula)

	// 3. database salva
	database.SalvarFuncionario(novoFuncionario)

	// 4. view exibe o resultado do cadastro
	views.ExibirFuncionario(novoFuncionario)
}

func ListarFuncionarios() {
	funcionarios := database.FuncionariosDB
	views.ExibirListaFuncionarios(funcionarios)
}

// função pra editar os dados do funcionário
func EditarFuncionario() {
	// 1. pede a identificação que é a chave pra achar o funcionário no banco
	matricula := views.SolicitarMatricula("editar")

	// 2. busca o funcionário no banco verificando se existe
	funcionario, existe := database.FuncionariosDB[matricula]
	if !existe {
		fmt.Print("Funcionário não encontrado.")
		return
	}

	// 3. exibe o funcionário que achou e solicita os novos dados
	views.ExibirFuncionario(funcionario)
	nome, login, telefone, email, cargo, matricula := views.SolicitarNovosDadosFuncionario(funcionario)

	// 4. atualiza os dados (lembrando que se o usuário apertar enter, mantém o valor que já tava)
	if nome != "" {
		funcionario.SetNome(nome)
	}
	if login != "" {
		funcionario.SetLogin(login)
	}
	if telefone != "" {
		funcionario.SetTelefone(telefone)
	}
	if email != "" {
		funcionario.SetEmail(email)
	}
	if cargo != "" {
		funcionario.SetCargo(cargo)
	}

	fmt.Print("Funcionário atualizado com sucesso!")
	views.ExibirFuncionario(funcionario)
}

// função pra deletar um funcionario
func DeletarFuncionario() {
	// 1. pede a matricula que é a chave pra achar o funcionario no banco
	matricula := views.SolicitarMatricula("deletar")

	// 2. busca o funcionario no banco verificando se existe
	funcionario, existe := database.FuncionariosDB[matricula]
	if !existe {
		fmt.Print("Funcionário não encontrado.")
		return
	}

	// 3. confirma se realmente quer deletar chamando a função da view
	if views.ConfirmarExclusaoFuncionario(funcionario) {
		// 4. se voltou como true, deleta o funcionario do banco
		delete(database.FuncionariosDB, matricula)
		fmt.Print("Funcionário excluído com sucesso.")
	} else {
		fmt.Print("Operação de exclusão cancelada.") //se não confirmou, cancela a exclusão
	}
}

func GerenciarFuncionarios() {
	for {
		views.MostrarSubMenu("Funcionários")
		opcao := utils.LerOpcao()

		switch opcao {
		case 1:
			CadastrarFuncionario()
		case 2:
			ListarFuncionarios()
		case 3:
			EditarFuncionario()
		case 4:
			DeletarFuncionario()
		case 0:
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
