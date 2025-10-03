package classes

import "fmt"

//
//
//Struct
type Produtos struct {
	nome              string
	descricao         string
	preco             float32
	quantidadeEstoque int
}

//
//
//Construtor
func AddProdutos() *Produtos {
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
	return &Produtos{nome: nome, descricao: descricao, preco: preco, quantidadeEstoque: quantidadeEstoque}
}

// Getters
func (item *Produtos) GetNome() string {
	return item.nome
}
func (item *Produtos) GetDescricao() string {
	return item.descricao
}
func (item *Produtos) GetPreco() float32 {
	return item.preco
}
func (item *Produtos) GetEstoque() int {
	return item.quantidadeEstoque
}

// Setters
func (item *Produtos) SetNome() {
	var temp string
	fmt.Println("Nome: ")
	fmt.Scan(temp)
	item.nome = temp
}
func (item *Produtos) SetDescricao(descricao string) {
	var temp string
	fmt.Println("Descrição: ")
	fmt.Scan(temp)
	item.descricao = descricao
}
func (item *Produtos) SetPreco() {
	var temp float32
	fmt.Println("Preço: ")
	fmt.Scan(temp)
	item.preco = temp
}
func (item *Produtos) SetEstoque() {
	var temp int
	fmt.Println("Estoque: ")
	fmt.Scan(temp)
	item.quantidadeEstoque = temp
}

/*Os Eventos Deletar, Consultar e Listar serão realizados no Banco de Dados*/
