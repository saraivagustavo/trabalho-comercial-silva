package models

import (
	"fmt"
	"time"
)

type Pagamento struct {
	efetuado bool
	forma    string
}

type NotaFiscal struct {
	pedido       *Pedido
	dataEfetuada time.Time
}

func NewNotaFiscal(pagamento *Pagamento, pedido *Pedido) *NotaFiscal {
	if pagamento.efetuado {
		return &NotaFiscal{pedido: pedido, dataEfetuada: time.Now()}
	} else {
		return nil
	}
}

func NewPagamento(forma string, efetuado bool) *Pagamento {
	return &Pagamento{efetuado: efetuado, forma: forma}
}
func (nf *NotaFiscal) ToString() string {
	if len(pedido.GetItens()) == 0 { //verifica se o cliente adicionou algum item
		fmt.Println("Nenhum item adicionado.")
	} else {
		for _, item := range pedido.GetItens() { //loop pra percorrer os itens do pedido
			fmt.Printf("  - %s | Quantidade: %d | Preço Unit.: R$ %.2f\n", //Printf pra formatar a string, similar ao printf do C ou print(f"") do Python
				item.GetProduto().GetNome(), //vai entrar no lugar do '%s'
				item.GetQuantidade(),        //vai entrar no lugar do '%d'
				item.GetPreco())             //vai entrar no lugar do '%.2f'
		}
	}

	return fmt.Sprintf(`
=========================================
          NOTA FISCAL - Nº %04d
=========================================
Data de Emissão: %s
-----------------------------------------
CLIENTE:
  Nome: %s
  Documento: %s
-----------------------------------------
ITENS DO PEDIDO:%s
-----------------------------------------
VALOR TOTAL: R$ %.2f
-----------------------------------------
FORMA DE PAGAMENTO: %s
=========================================
`)
}
