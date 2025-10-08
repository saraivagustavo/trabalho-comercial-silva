package views

import (
	"bufio"
	"fmt"
	"os"
	"trabalho/models"
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

// ExibirListaProdutos pra listar todos os produtos cadastrados, vai ser bom pra mostrar pro cliente na hora de fazer o pedido o que tem disponível
func ExibirListaProdutos(produtos map[string]*models.Produto) {
	fmt.Println("\n--- LISTA DE PRODUTOS DISPONÍVEIS ---")
	if len(produtos) == 0 {
		fmt.Println("Nenhum produto cadastrado.")
		return
	}
	for _, produto := range produtos {
		fmt.Printf("- Nome: %s | Preço: R$ %.2f | Estoque: %d\n",
			produto.GetNome(),
			produto.GetPrecoProduto(),
			produto.GetQuantidadeEstoque())
	}
	fmt.Println("---------------------------------------")
}
