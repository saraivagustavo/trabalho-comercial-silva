package models

import (
	"errors" // Importa o pacote errors para criar erros personalizados (bom fazer isso nas outras classes do model também)
	"fmt"
)

type Produto struct {
	id                string
	nome              string
	descricao         string
	preco             float64
	quantidadeEstoque int
}

// construtor do Produto, os atributos são passados no parâmetro
func NewProduto(nome, descricao string, preco float64, quantidade int, id string) (*Produto, error) {
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
		id:                GerarIDProduto(),
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

func (p *Produto) GetID() string {
	return p.id
}

// ===== Setters da classe =====
func (p *Produto) SetNome(novoNome string) {
	p.nome = novoNome
}

func (p *Produto) SetDescricao(novaDescricao string) {
	p.descricao = novaDescricao
}

func (p *Produto) SetPreco(novoPreco float64) {
	p.preco = novoPreco
}

func (p *Produto) SetQuantidadeEstoque(novaQuantidade int) error {
	if novaQuantidade < 0 {
		return errors.New("o estoque não pode se tornar negativo")
	}
	p.quantidadeEstoque = novaQuantidade
	return nil
}

// GerarIDProduto gera um ID único pro produto, esse ID é usado como chave no map do banco de dados
var nextID = 1

func GerarIDProduto() string {
	id := fmt.Sprintf("PROD%d", nextID)
	nextID++
	return id
}

// ToString retorna uma representação em string do produto
func (p *Produto) ToString() string {
	return "\n--- DADOS DO PRODUTO ---\n" +
		"ID: " + p.id + "\n" +
		"Nome: " + p.nome + "\n" +
		"Descrição: " + p.descricao + "\n" +
		fmt.Sprintf("Preço: R$ %.2f\n", p.preco) +
		fmt.Sprintf("Quantidade em Estoque: %d\n", p.quantidadeEstoque) +
		"------------------------"
}
