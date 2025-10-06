package views

import (
	"bufio"
	"fmt"
	"os"
	"trabalho/models"
)

//mesma lógica do cliente_view
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
//mesma lógica do cliente_view
func ExibirFuncionario(funcionario *models.Funcionario) {
	fmt.Println("\n--- FUNCIONÁRIO CADASTRADO ---")
	fmt.Println(funcionario.ToString())
	fmt.Println("------------------------------")
}