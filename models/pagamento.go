package models

import (
	"fmt"
	"time"
	"trabalho/utils"
)

type Pagamento struct {
	pedido   *Pedido
	time     string
	forma    string
	aprovado bool
}

func NewPagamento(forma string, pedido *Pedido, aprovado bool) *Pagamento {
	return &Pagamento{forma: forma, pedido: pedido, time: utils.TimeToString(time.Now()), aprovado: aprovado}
}

// Set
func (pag *Pagamento) SetAprovado(status bool) {
	pag.aprovado = status
}

// Getters
func (pag *Pagamento) GetData() string {
	return pag.time
}
func (pag *Pagamento) GetPedido() *Pedido {
	return pag.pedido
}
func (pag *Pagamento) GetForma() string {
	return pag.forma
}

func (nota *Pagamento) ToString() string {
	var pedidosNota string = ""
	if len(nota.pedido.GetItens()) == 0 { //verifica se o cliente adicionou algum item
		pedidosNota += "Nenhum item adicionado."
	} else {
		for _, item := range nota.pedido.GetItens() { //loop pra percorrer os itens do pedido
			pedidosNota += fmt.Sprintf("  - %s | Quantidade: %d | Preço Unit.: R$ %.2f\n", //Printf pra formatar a string, similar ao printf do C ou print(f"") do Python
				item.GetProduto().GetNome(), //vai entrar no lugar do '%s'
				item.GetQuantidade(),        //vai entrar no lugar do '%d'
				item.GetPreco())             //vai entrar no lugar do '%.2f'
		}
	}

	return fmt.Sprintf(`
 =========================================
               NOTA FISCAL
 =========================================
 Data de Emissão: %s
 -----------------------------------------
 CLIENTE:
   Nome: %s
   Documento: %s
 -----------------------------------------
 ITENS DO PEDIDO: 
 %s
 -----------------------------------------
 VALOR TOTAL: R$ %.2f
 -----------------------------------------
 FORMA DE PAGAMENTO: %s
 =========================================
 `, nota.GetData(), nota.GetPedido().GetCliente().GetNome(), nota.GetPedido().GetCliente().GetIdentificacao(), pedidosNota, nota.GetPedido().GetValorTotal(), nota.GetForma())
}
