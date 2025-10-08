package views

import (
	"fmt"
	"trabalho/models"
	"trabalho/utils"
)

func FazerPagamento() *models.Pagamento {
	utils.clearScreen()
	fmt.Println("Escolha sua forma de pagamento:\n 1.Pix\n 2. Boleto\n 3.Cartão \n")
	resposta := utils.LerString()

}
