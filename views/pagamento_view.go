package views

import (
	"fmt"
	"trabalho/utils"
)

func SelecionarFormaPagamento() int {
	fmt.Println("-------------------------------")
	fmt.Println("SELECIONE A FORMA DE PAGAMENTO:")
	fmt.Println("1 - Pix")
	fmt.Println("2 - Cartão")
	fmt.Println("3 - Boleto")
	fmt.Println("0 - Cancelar")
	return utils.LerOpcao()
}

func ExibirNota(str string) {
	fmt.Printf("%s", str)
}
