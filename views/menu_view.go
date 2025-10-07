package views

import "fmt"

func MostrarMenu() {
	fmt.Println("\n====== COMERCIAL SILVA ======")
	fmt.Println("1. Cadastrar Cliente")
	fmt.Println("2. Cadastrar Funcionário")
	fmt.Println("0. Sair")
	fmt.Println("=============================")
	fmt.Print("Opção: ")
}
