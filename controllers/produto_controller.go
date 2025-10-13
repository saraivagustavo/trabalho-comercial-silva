package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"trabalho/database"
	"trabalho/models"
	"trabalho/utils"
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
		fmt.Print("Preço inválido.")
		return
	}

	quantidade, errQntd := strconv.Atoi(strings.TrimSpace(quantidadeStr)) //conversão da quantidade em estoque
	if errQntd != nil {                                                   //mesma coisa, verifica e se der erro é pq o valor digitado não é válido
		fmt.Print("Quantidade inválida.")
		return
	}

	// 3. nesse passo que o produto é criado, o controller chama o construtor que tá no model e passa os dados que puxou da view (linha 16)
	id := models.GerarIDProduto() //gera o ID único pro produto
	novoProduto, err := models.NewProduto(nome, descricao, preco, quantidade, id)
	if err != nil {
		fmt.Print("Erro ao criar produto: " + err.Error())
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

// função pra editar os dados do produto
func EditarProduto() {
	// 1. pede o id que é a chave pra achar o produto no banco
	id := views.SolicitarIDProduto()

	// 2. busca o id no banco verificando se existe
	produto, existe := database.ProdutosDB[id]
	if !existe {
		fmt.Print("Produto não encontrado.")
		return
	}

	// 3. exibe o produto que achou e solicita os novos dados
	views.ExibirProduto(produto)
	nome, descricao, precoStr, quantidadeStr := views.SolicitarNovosDadosProduto(produto)

	// 4. atualiza os dados (lembrando que se o usuário apertar enter, mantém o valor que já tava)
	if nome != "" {
		produto.SetNome(nome)
	}
	if descricao != "" {
		produto.SetDescricao(descricao)
	}
	if precoStr != "" {
		novoPreco, err := strconv.ParseFloat(precoStr, 64)
		if err != nil {
			fmt.Print("Preço inválido.")
		} else {
			produto.SetPreco(novoPreco)
		}
	}
	if quantidadeStr != "" {
		novaQuantidade, err := strconv.Atoi(quantidadeStr)
		if err != nil {
			fmt.Print("Quantidade inválida.")
		} else if err = produto.SetQuantidadeEstoque(novaQuantidade); err != nil {
			fmt.Print("ERRO: " + err.Error())
		}
	}

	fmt.Print("Funcionário atualizado com sucesso!")
	views.ExibirProduto(produto)
}

// função pra deletar um produto
func DeletarProduto() {
	// 1. pede o id que é a chave pra achar o produto no banco
	id := views.SolicitarIDProduto()

	// 2. busca o produto no banco verificando se existe
	produto, existe := database.ProdutosDB[id]
	if !existe {
		fmt.Print("Produto não encontrado.")
		return
	}

	// 3. confirma se realmente quer deletar chamando a função da view
	if views.ConfirmarExclusaoProduto(produto) {
		// 4. se voltou como true, deleta o produto do banco
		delete(database.ProdutosDB, id)
		fmt.Print("Produto excluído com sucesso.")
	} else {
		fmt.Print("Operação de exclusão cancelada.") //se não confirmou, cancela a exclusão
	}
}

func GerenciarProdutos() {
	for {
		views.MostrarSubMenu("Produtos")
		opcao := utils.LerOpcao()

		switch opcao {
		case 1:
			CadastrarProduto()
		case 2:
			ListarProdutos()
		case 3:
			EditarProduto()
		case 4:
			DeletarProduto()
		case 0:
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
