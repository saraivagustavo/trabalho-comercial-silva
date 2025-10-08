package main

import (
	"fmt"
	"trabalho/controllers"
	"trabalho/utils"
	"trabalho/views"
)

func main() {
	var opcao int

	for {
		views.MostrarMenuPrincipal()
		opcao = utils.LerOpcao()

		switch opcao {
		case 1:
			controllers.GerenciarClientes()

		case 2:
			controllers.GerenciarFuncionarios()

		case 3:
			controllers.GerenciarProdutos()

		case 4:
			controllers.GerenciarPedidos()

		case 0:
			fmt.Println("\nSaindo...")
			return

		default:
			fmt.Println("\nOpção inválida, tente novamente!")

		}
	}

}
