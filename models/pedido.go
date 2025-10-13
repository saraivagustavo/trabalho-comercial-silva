package models

import (
	"errors"
	"time"
)

// ItemPedido vai representar cada produto dentro de um único pedido
type ItemPedido struct {
	produto    *Produto
	quantidade int
	preco      float64
}

// representa a solicitação total de um cliente com todos os itens do pedido
type Pedido struct {
	cliente    *Cliente
	itens      []*ItemPedido
	valorTotal float64
	data       time.Time
	confirmado bool
}

// Getters pro ItemPedido
func (item_pedido *ItemPedido) GetProduto() *Produto {
	return item_pedido.produto
}

func (item_pedido *ItemPedido) GetQuantidade() int {
	return item_pedido.quantidade
}

func (item_pedido *ItemPedido) GetPreco() float64 {
	return item_pedido.preco
}

// Getters pro Pedido
func (pedido *Pedido) GetCliente() *Cliente {
	return pedido.cliente
}

func (pedido *Pedido) GetItens() []*ItemPedido {
	return pedido.itens
}

func (pedido *Pedido) GetValorTotal() float64 {
	return pedido.valorTotal
}

func (pedido *Pedido) GetData() time.Time {
	return pedido.data
}

func (pedido *Pedido) GetConfirmado() bool {
	return pedido.confirmado
}

//Setter que eu senti falta ;-;

func (pedido *Pedido) SetConfirmado(bo bool) {
	pedido.confirmado = bo
}

// Construtor pro Pedido
func NewPedido(cliente *Cliente) *Pedido {
	return &Pedido{
		cliente:    cliente,
		itens:      []*ItemPedido{},
		data:       time.Now(),
		confirmado: false,
	}
}

// AdicionarItem adiciona um produto ao pedido, verificando antes, a quantidade no estoque
func (p *Pedido) AdicionarItem(produto *Produto, quantidade int) error {
	if produto.GetQuantidadeEstoque() < quantidade {
		return errors.New("estoque insuficiente")
	}

	item := &ItemPedido{
		produto:    produto,
		quantidade: quantidade,
		preco:      float64(produto.GetPrecoProduto()), // Usamos o Getter de Produto
	}
	p.itens = append(p.itens, item)

	produto.SetQuantidadeEstoque(produto.GetQuantidadeEstoque() - quantidade) // Atualiza o estoque do produto

	p.calcularValorTotal()
	return nil
}

// calcularValorTotal atualiza o valor total do pedido
func (p *Pedido) calcularValorTotal() {
	var total float64
	for _, item := range p.itens {
		// usamos os getters do item
		total += item.GetPreco() * float64(item.GetQuantidade())
	}
	p.valorTotal = total
}

// finalizar o produto e confirmar o pedido
func (p *Pedido) Confirmar() {
	p.confirmado = true
}
