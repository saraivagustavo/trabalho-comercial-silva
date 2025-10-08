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

// CriarPedidoController gerencia o fluxo de criação de um novo pedido
func CriarPedidoController() {
	//o controller solicita a identificação do cliente pra view
	identificacao := views.SolicitarIdentificacaoCliente()
	identificacao = strings.TrimSpace(identificacao) //faz a limpeza do \n

	//o controller busca o cliente no banco de dados usando a identificação que puxou da view
	cliente, existe := database.ClientesDB[identificacao]
	if !existe {
		fmt.Print("Cliente com documento '" + identificacao + "' não encontrado!")
		return
	}
	fmt.Print("Cliente encontrado: " + cliente.GetNome())

	//o controller cria um novo pedido pro cliente usando o model dele
	novoPedido := models.NewPedido(cliente)

	for {
		//só pra teste, exibe os detalhes do pedido a cada iteração
		views.ExibirDetalhesPedido(novoPedido)
		views.ExibirListaProdutos(database.ProdutosDB)

		//solicita o produto e a quantidade do produto pra view
		idProduto, quantidadeStr := views.SelecionarProduto()

		//se for vazia, saí do loop
		if idProduto == "" {
			break
		}

		//limpa a string da quantidade
		quantidadeStr = strings.TrimSpace(quantidadeStr)
		//depois de limpar, tem que converter a string pra int
		quantidade, err := strconv.Atoi(quantidadeStr)
		if err != nil || quantidade <= 0 {
			fmt.Print("Quantidade inválida. Por favor, digite um número inteiro maior que zero.")
			continue
		}

		// o controller tem que verificar se o produto existe no banco de dados
		produto, existe := database.ProdutosDB[idProduto]
		if !existe {
			fmt.Print("Produto '" + idProduto + "' não encontrado!")
			continue
		}

		//adiciona o item ao pedido usando o model dele
		novoPedido.AdicionarItem(produto, quantidade)
		fmt.Printf("Adicionado %d x %s ao pedido.\n", quantidade, produto.GetNome())
	}

	// verifica se o pedido tem itens antes de confirmar e chamar a função de confirmação
	if len(novoPedido.GetItens()) == 0 {
		fmt.Print("Nenhum item foi adicionado. Pedido cancelado.")
		return
	}

	views.ExibirDetalhesPedido(novoPedido)
	if views.ConfirmarPedido() { // se a view retornar true, confirma o pedido
		novoPedido.Confirmar()                        //altera o status do pedido pra confirmado
		idPedido := database.SalvarPedido(novoPedido) //salva o pedido no banco e pega o ID retornado

		views.ExibirConfirmacaoPedido(novoPedido) //chama a função da view pra exibir a confirmação do pedido

		fmt.Printf("ID do Pedido: #%d\n", idPedido) //exibe o ID do pedido (depois que salvou no banco, incrementa o ID pro próximo pedido)

	} else { //se a view retornar false, cancela o pedido
		fmt.Print("Pedido cancelado pelo usuário.")
	}
}

// função pra listar os pedidos do banco
func ListarPedidos() {
	pedidos := database.PedidosDB
	if len(pedidos) == 0 {
		fmt.Print("Nenhum pedido foi realizado ainda.")
		return
	}
	for id, pedido := range pedidos {
		fmt.Printf("\n--- PEDIDO ID: #%d ---\n", id)
		views.ExibirDetalhesPedido(pedido)
	}
}

func GerenciarPedidos() {
	for {
		views.MostrarSubMenu("Pedidos")
		opcao := utils.LerOpcao()

		switch opcao {
		case 1:
			CriarPedidoController()
		case 2:
			ListarPedidos()
		case 3:

		case 4:
			fmt.Print("Funcionalidade ainda não implementada.")
		case 0:
			return
		default:
			fmt.Print("Opção inválida!")
		}
	}
}
