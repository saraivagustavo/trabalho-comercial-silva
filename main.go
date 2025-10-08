package main

import (
	"fmt"
	"trabalho/controllers"
	"trabalho/views"
)

func main() {
	var opcao int

	for {
		views.MostrarMenu()
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			fmt.Println("\nCadastrando Cliente...")
			controllers.CadastrarCliente()

		case 2:
			fmt.Println("\nListando Clientes...")
			controllers.ListarClientes()

		case 3:
			fmt.Println("\nCadastrando Funcionário...")
			controllers.CadastrarFuncionario()

		case 4:
			fmt.Println("\nListando Funcionários...")
			controllers.ListarFuncionarios()

		case 0:
			fmt.Println("\nSaindo...")
			return

		default:
			fmt.Println("\nOpção inválida, tente novamente!")

		}
	}

}
