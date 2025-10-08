package models

import "errors" // Importa o pacote errors para criar erros personalizados (bom fazer isso nas outras classes do model também)

type Produto struct {
	nome              string
	descricao         string
	preco             float64
	quantidadeEstoque int
}

//construtor do Produto, os atributos são passados no parâmetro
func NewProduto(nome, descricao string, preco float64, quantidade int) (*Produto, error) {
	//faz as validações do minimundo do produto
	if preco <= 0 {
		return nil, errors.New("o preço do produto não pode ser negativo ou igual a zero")
	}
	if quantidade < 0 {
		return nil, errors.New("o estoque do produto não pode ser negativo")
	}

	return &Produto{
		nome:              nome,
		descricao:         descricao,
		preco:             preco,
		quantidadeEstoque: quantidade,
	}, nil
}

// ===== Getters da classe =====
func (p *Produto) GetNome() string {
	return p.nome
}

func (p *Produto) GetDescricao() string {
	return p.descricao
}

func (p *Produto) GetPrecoProduto() float64 {
	return p.preco
}

func (p *Produto) GetQuantidadeEstoque() int {
	return p.quantidadeEstoque
}

// ===== Setters da classe ===== (aparentemente só precisa do setter do estoque)
func (p *Produto) SetQuantidadeEstoque(novaQuantidade int) error {
	if novaQuantidade < 0 {
		return errors.New("o estoque não pode se tornar negativo")
	}
	p.quantidadeEstoque = novaQuantidade
	return nil
}
