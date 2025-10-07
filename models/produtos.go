package models

import "fmt"

//
//
//Struct
type Produto struct {
	nome              string
	descricao         string
	preco             float32
	quantidadeEstoque int
}

//
//
//Construtor
func AddProdutos() *Produto {
	var nome string
	var descricao string
	var preco float32
	var quantidadeEstoque int

	fmt.Println("Nome: ")
	fmt.Scan(nome)
	fmt.Println("Descrição: ")
	fmt.Scan(descricao)
	fmt.Println("Preço: ")
	fmt.Scan(preco)
	fmt.Println("Estoque: ")
	fmt.Scan(quantidadeEstoque)

	if preco <= 0 {
		fmt.Println("O preço do produto não pode ser negativo ou igual a zero")
		AddProdutos()
	}
	if quantidadeEstoque < 0 {
		fmt.Println("O estoque do produto não pode ser negativo")
		AddProdutos()
	}
	return &Produto{nome: nome, descricao: descricao, preco: preco, quantidadeEstoque: quantidadeEstoque}
}

// Getters
func (item *Produto) GetNome() string {
	return item.nome
}
func (item *Produto) GetDescricao() string {
	return item.descricao
}
func (item *Produto) GetPrecoProduto() float32 {
	return item.preco
}
func (item *Produto) GetQuantidadeEstoque() int {
	return item.quantidadeEstoque
}

// Setters
func (item *Produto) SetNome() {
	var temp string
	fmt.Println("Nome: ")
	fmt.Scan(temp)
	item.nome = temp
}
func (item *Produto) SetDescricao(descricao string) {
	var temp string
	fmt.Println("Descrição: ")
	fmt.Scan(temp)
	item.descricao = descricao
}
func (item *Produto) SetPreco() {
	var temp float32
	fmt.Println("Preço: ")
	fmt.Scan(temp)
	item.preco = temp
}
func (item *Produto) SetQuantidadeEstoque() {
	var temp int
	fmt.Println("Estoque: ")
	fmt.Scan(temp)
	item.quantidadeEstoque = temp
}

/*Os Eventos Deletar, Consultar e Listar serão realizados no Banco de Dados*/
