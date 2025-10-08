package views

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"trabalho/models"
	"trabalho/utils"
)

// SolicitarDadosProduto pede os dados para cadastrar um novo produto
func SolicitarDadosProduto() (nome, descricao, precoStr, quantidadeStr string) {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- CADASTRO DE NOVO PRODUTO ---")
	fmt.Print("Nome do Produto: ")
	nome, _ = scanner.ReadString('\n')

	fmt.Print("Descrição: ")
	descricao, _ = scanner.ReadString('\n')

	fmt.Print("Preço (ex: 150.99): ")
	precoStr, _ = scanner.ReadString('\n')

	fmt.Print("Quantidade em Estoque: ")
	quantidadeStr, _ = scanner.ReadString('\n')

	return
}

// ExibirProduto exibe os dados de um produto passado como parâmetro
func ExibirProduto(produto *models.Produto) {
	fmt.Println("\n--- PRODUTO ---")
	fmt.Println(produto.ToString())
	fmt.Println("----------------")
}

// ExibirListaProdutos pra listar todos os produtos cadastrados, vai ser bom pra mostrar pro cliente na hora de fazer o pedido o que tem disponível
func ExibirListaProdutos(produtos map[string]*models.Produto) {
	fmt.Println("\n--- LISTA DE PRODUTOS DISPONÍVEIS ---")
	if len(produtos) == 0 {
		fmt.Println("Nenhum produto cadastrado.")
		return
	}
	for _, produto := range produtos {
		fmt.Printf("- ID: %s | Nome: %s | Preço: R$ %.2f | Estoque: %d\n",
			produto.GetID(),
			produto.GetNome(),
			produto.GetPrecoProduto(),
			produto.GetQuantidadeEstoque())
	}
	fmt.Println("---------------------------------------")
}

func SolicitarIDProduto() string {
	fmt.Print("Digite o id do produto: ")
	id := utils.LerString()
	return id
}

// SolicitarNovosDadosProduto pra atualizar os dados do produto, se não quiser mudar todos os dados, se o usuário apertar enter, mantém o valor que já tava
func SolicitarNovosDadosProduto(produto *models.Produto) (nome, descricao, precoStr, quantidadeStr string) {
	fmt.Println("\nDigite os novos dados. Pressione ENTER para manter o valor atual.")

	// nome
	fmt.Printf("Nome atual: %s\nNovo nome: ", produto.GetNome())
	nome = utils.LerString()

	// descrição
	fmt.Printf("Descrição atual: %s\nNova descrição: ", produto.GetDescricao())
	descricao = utils.LerString()

	// preço
	fmt.Printf("Preço atual: R$ %.2f\nNovo preço: ", produto.GetPrecoProduto())
	precoStr = utils.LerString()

	// quantidade
	fmt.Printf("Quantidade atual: %d\nNova quantidade: ", produto.GetQuantidadeEstoque())
	quantidadeStr = utils.LerString()

	return
}

// ConfirmarExclusao só pra segurança de não deletar o produto errado sem querer
func ConfirmarExclusaoProduto(produto *models.Produto) bool {
	fmt.Printf("Tem certeza que deseja excluir o produto: %s\n", produto.GetNome())
	fmt.Print("Digite 's' para confirmar: ")
	resposta := utils.LerString()
	return strings.ToLower(resposta) == "s" //se for 's' isso da true e confirma a exclusão
}
