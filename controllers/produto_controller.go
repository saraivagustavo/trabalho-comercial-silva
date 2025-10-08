// Arquivo: controllers/produto_controller.go
package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"trabalho/database"
	"trabalho/models"
	"trabalho/views"
)

// função que gerencia o cadastro de um novo produto
func CadastrarProduto() {
	// 1. o controller chama a view para obter os que o usuário digitou
	nome, descricao, precoStr, quantidadeStr := views.SolicitarDadosProduto()

	// 2. controller limpa, converte e valida os dados
	nome = strings.TrimSpace(nome)           //limpeza do nome
	descricao = strings.TrimSpace(descricao) //limpeza da descrição

	preco, errPreco := strconv.ParseFloat(strings.TrimSpace(precoStr), 64) //conversão do preço
	if errPreco != nil {                                                   //verificação de erro na conversão, se der erro é pq o valor digitado não é válido
		fmt.Print("ERRO: Preço inválido.")
		return
	}

	quantidade, errQtde := strconv.Atoi(strings.TrimSpace(quantidadeStr)) //conversão da quantidade em estoque
	if errQtde != nil {                                                   //mesma coisa, verifica e se der erro é pq o valor digitado não é válido
		fmt.Print("ERRO: Quantidade inválida.")
		return
	}

	// 3. nesse passo que o produto é criado, o controller chama o construtor que tá no model e passa os dados que puxou da view (linha 16)
	novoProduto, err := models.NewProduto(nome, descricao, preco, quantidade)
	if err != nil {
		fmt.Print("ERRO ao criar produto: " + err.Error())
		return
	}

	// 4. aqui salva o produto criado no banco, só precisa chamar a função que tá no database
	database.SalvarProduto(novoProduto)

	// 5. confirmação que foi cadastrado
	fmt.Printf("Produto '%s' cadastrado com sucesso!", novoProduto.GetNome())
}

// função pra listar os produtos cadastrados
func ListarProdutos() {
	produtos := database.ProdutosDB     //puxa do banco os produtos cadastrados
	views.ExibirListaProdutos(produtos) //chama a a função da view que exibe a lista
}
