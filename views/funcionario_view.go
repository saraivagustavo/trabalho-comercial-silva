package views

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"trabalho/models"
	"trabalho/utils"
)

// mesma lógica do cliente_view
func SolicitarDadosFuncionario() (nome, login, tipo, telefone, email, cargo, matricula string) {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- CADASTRO DE NOVO FUNCIONÁRIO ---")

	fmt.Print("Nome Completo: ")
	nome, _ = scanner.ReadString('\n')

	fmt.Print("Login: ")
	login, _ = scanner.ReadString('\n')

	fmt.Print("Tipo/Perfil: ")
	tipo, _ = scanner.ReadString('\n')

	fmt.Print("Telefone (apenas números, 11 dígitos): ")
	telefone, _ = scanner.ReadString('\n')

	fmt.Print("Email: ")
	email, _ = scanner.ReadString('\n')

	fmt.Print("Cargo: ")
	cargo, _ = scanner.ReadString('\n')

	fmt.Print("Matrícula: ")
	matricula, _ = scanner.ReadString('\n')
	fmt.Println("-------------------------------")

	return
}

// mesma lógica do cliente_view
func ExibirFuncionario(funcionario *models.Funcionario) {
	fmt.Println("\n--- FUNCIONÁRIO CADASTRADO ---")
	fmt.Println(funcionario.ToString())
	fmt.Println("------------------------------")
}

// segue a mesma lógica do cliente_view pra exibir a lista de funcionários
func ExibirListaFuncionarios(funcionarios map[string]*models.Funcionario) {
	fmt.Println("\n--- LISTA DE FUNCIONÁRIOS ---")
	if len(funcionarios) == 0 {
		fmt.Println("Nenhum funcionário cadastrado.")
		return
	}
	for _, funcionario := range funcionarios {
		fmt.Println(funcionario.ToString())
	}
}

// SolicitarMatricula segue a mesma lógica do cliente_view pra pedir a identificação do funcionário, mas aqui é a matrícula
func SolicitarMatricula(acao string) string {
	fmt.Printf("\n--- %s Funcionário ---\n", acao) //ação vai ser "Editar" ou "Deletar", vai puxar a string que foi passada como parâmetro
	fmt.Print("Digite a matrícula do funcionário: ")
	return utils.LerString()
}

// SolicitarNovosDadosFuncionario pra atualizar os dados do funcionário, se não quiser mudar todos os dados, se o usuário apertar enter, mantém o valor que já tava
func SolicitarNovosDadosFuncionario(funcionario *models.Funcionario) (nome, login, telefone, email, cargo, matricula string) {
	fmt.Println("\nDigite os novos dados. Pressione ENTER para manter o valor atual.")

	// nome
	fmt.Printf("Nome atual: %s\nNovo nome: ", funcionario.GetNome())
	nome = utils.LerString()

	// login
	fmt.Printf("Login atual: %s\nNovo login: ", funcionario.GetLogin())
	login = utils.LerString()

	// telefone
	fmt.Printf("Telefone atual: %s\nNovo telefone: ", funcionario.GetTelefone())
	telefone = utils.LerString()

	// email
	fmt.Printf("Email atual: %s\nNovo email: ", funcionario.GetEmail())
	email = utils.LerString()

	// cargo
	fmt.Printf("Cargo atual: %s\nNovo cargo: ", funcionario.GetCargo())
	cargo = utils.LerString()

	return
}

// ConfirmarExclusao só pra segurança de não deletar o funcionário errado sem querer
func ConfirmarExclusaoFuncionario(funcionario *models.Funcionario) bool {
	fmt.Printf("Tem certeza que deseja excluir o funcionário: %s (Matrícula: %s)?\n", funcionario.GetNome(), funcionario.GetMatricula())
	fmt.Print("Digite 's' para confirmar: ")
	resposta := utils.LerString()
	return strings.ToLower(resposta) == "s" //se for 's' isso da true e confirma a exclusão
}
