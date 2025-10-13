package controllers

import (
	"fmt"
	"math/rand"
	"time"
	"trabalho/models"
	"trabalho/utils"
	"trabalho/views"
)

func ProcessarPagamento(pedido models.Pedido) {
	var ans int
	var forma string
	var aprovado bool
	rand.Seed(time.Now().UnixNano())
loop:
	for {
		ans = views.SelecionarFormaPagamento()
		switch ans {
		case 0:
			fmt.Println("Pagamento Cancelado...")
			return
		case 1:
			forma = "Pix"
			aprovado = true
			break loop
		case 2:
			forma = "Cartão"
			aprovado = rand.Intn(2) == 0
			break loop
		case 3:
			forma = "Boleto"
			aprovado = rand.Intn(3) == 0
			break loop
		default:
			fmt.Println("Opção Inválida. Tente novamente")
			continue
		}
	}

	if aprovado {
		pag := models.NewPagamento(forma, pedido).ToString()
		views.ExibirNota(pag)
		pedido.SetConfirmado(true)
		views.MostrarMenuPrincipal()
	} else {
		for {
			fmt.Println("O pagamento foi recusado. Deseja tentar novamente? (s/n)")
			resp := utils.LerString()
			switch resp {
			case "s":
				ProcessarPagamento(pedido)
			case "n":
				views.MostrarMenuPrincipal()
			}
		}
	}
}
