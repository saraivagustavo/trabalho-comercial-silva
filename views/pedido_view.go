package views

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"trabalho/models"
	"trabalho/utils"
	//"time"
)

// função pra solicitar a identificação do cliente pra iniciar um novo pedido
func SolicitarIdentificacaoCliente() string {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- INICIARUM NOVO PEDIDO ---")
	fmt.Print("Informe o CPF/CNPJ do cliente: ")
	identificacao, _ := scanner.ReadString('\n')
	return identificacao
}

// função pra solicitar ao usuário os produtos que ele vai querer no pedido
func SelecionarProduto() (idProduto string, quantidadeStr string) {
	fmt.Print("Digite o id do produto que deseja adicionar (ou 'fim' para sair): ")
	idProduto = utils.LerString() //lê o id do produto

	if strings.TrimSpace(idProduto) != "" {
		fmt.Printf("Digite a quantidade de '%s': ", strings.TrimSpace(idProduto))
		quantidadeStr = utils.LerString() // Lê a quantidade como string
	}
	return idProduto, quantidadeStr //retorna o id do produto e a quantidade
}

// função pra exibir o resumo do pedido do cliente
func ExibirDetalhesPedido(pedido *models.Pedido) {
	fmt.Println("\n--- RESUMO DO PEDIDO ---")
	fmt.Printf("Cliente: %s\n", pedido.GetCliente().GetNome()) //usa o getter do cliente pra pegar o nome do cliente
	fmt.Println("Itens: ")
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
	fmt.Printf("Valor Total: R$ %.2f\n", pedido.GetValorTotal()) //exibe o valor total do pedido usando o getter do pedido que já calcula o valor total
}

// função pra confirmar o pedido e alterar o status de confirmado pra true
func ConfirmarPedido() bool {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Deseja confirmar o pedido? (s/n): ")
	resposta, _ := scanner.ReadString('\n')
	resposta = strings.ToLower(strings.TrimSpace(resposta))
	return resposta == "s" || resposta == "sim"
}

// função só pra exibir a confirmação do pedido
func ExibirConfirmacaoPedido(pedido *models.Pedido) {
	fmt.Println("\n--- PEDIDO CONFIRMADO ---")
	fmt.Printf("Pedido para o cliente %s confirmado com sucesso!\n", pedido.GetCliente().GetNome())
	fmt.Printf("Data do Pedido: %s\n", pedido.GetData())
	fmt.Printf("Valor Total: R$ %.2f\n", pedido.GetValorTotal())
	fmt.Println("-------------------------")
}
