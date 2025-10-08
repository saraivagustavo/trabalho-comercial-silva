package database

import "trabalho/models"

// ClientesDB é o banco pra armazenar os clientes
// A chave do mapa vai ser o cpf ou cnpj do cliente
var ClientesDB = make(map[string]*models.Cliente) //usar map aqui pq é mais fácil de buscar depois

//mesma coisa, mas pros funcionários e com a chaave sendo a matrícula dele
var FuncionariosDB = make(map[string]*models.Funcionario)

//mesma coisa pra armazenar os produtos, a chave é o nome do produto
var ProdutosDB = make(map[string]*models.Produto)

// PedidosDB armazena os pedidos, a chave é um ID incremental
var PedidosDB = make(map[int]*models.Pedido)
var proximoIDPedido = 1 //inicia um ID pra cada pedido que for criado pra ser usado como a chave do map

//a função SalvarCliente adiciona um novo cliente no banco
func SalvarCliente(cliente *models.Cliente) {
	ClientesDB[cliente.GetIdentificacao()] = cliente
}

//" " " "
func SalvarFuncionario(funcionario *models.Funcionario) {
	FuncionariosDB[funcionario.GetMatricula()] = funcionario
}

// SalvarProduto adiciona um novo produto
func SalvarProduto(produto *models.Produto) {
	ProdutosDB[produto.GetNome()] = produto
}

// SalvarPedido adiciona um novo pedido e retorna o ID do pedido e incrementa o ID pro próximo pedido
func SalvarPedido(pedido *models.Pedido) int {
	id := proximoIDPedido
	PedidosDB[id] = pedido
	proximoIDPedido++
	return id
}
