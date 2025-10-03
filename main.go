package main

import (
	"fmt"
	"trabalho/classes"
)

func main() {
	cliente := classes.NewCliente()

	funcionario := classes.NewFuncionario()

	fmt.Println(cliente.ToString())
	fmt.Println(funcionario.ToString())

}
