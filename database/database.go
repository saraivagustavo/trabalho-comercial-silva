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
	ProdutosDB[produto.GetID()] = produto
}

// SalvarPedido adiciona um novo pedido e retorna o ID do pedido e incrementa o ID pro próximo pedido
func SalvarPedido(pedido *models.Pedido) int {
	id := proximoIDPedido
	PedidosDB[id] = pedido
	proximoIDPedido++
	return id
}

//carga de dados que gpt fez pra testar o sistema
func init() {
	// --- CARGA DE PRODUTOS ---
	p1, _ := models.NewProduto("Notebook Dell", "Core i7, 16GB RAM", 4500.00, 10, "")
	p2, _ := models.NewProduto("Mouse Logitech", "Mouse sem fio ergonômico", 150.50, 50, "")
	p3, _ := models.NewProduto("Teclado Mecanico", "Teclado Redragon RGB", 350.00, 30, "")
	p4, _ := models.NewProduto("Monitor 24 polegadas", "Monitor LG Full HD 75Hz", 950.99, 15, "")
	p5, _ := models.NewProduto("SSD 1TB", "SSD Kingston NVMe M.2", 650.00, 40, "")
	SalvarProduto(p1)
	SalvarProduto(p2)
	SalvarProduto(p3)
	SalvarProduto(p4)
	SalvarProduto(p5)

	// --- CARGA DE CLIENTES ---
	c1 := models.NewCliente("Ana Silva", "ana.silva", "cliente", "11987654321", "ana.silva@email.com", "11122233344", "Rua das Flores, 123")
	c2 := models.NewCliente("Bruno Costa", "bruno.costa", "cliente", "21912345678", "bruno.costa@email.com", "99988877766", "Avenida Central, 456")
	c3 := models.NewCliente("Carla Souza", "carla.souza", "cliente", "31998761234", "carla.souza@email.com", "12345678901", "Praça da Matriz, 789")
	SalvarCliente(c1)
	SalvarCliente(c2)
	SalvarCliente(c3)

	// --- CARGA DE FUNCIONÁRIOS ---
	f1 := models.NewFuncionario("Carlos Oliveira", "carlos.o", "123", "vendedor", "11955554444", "carlos.o@comercialsilva.com", "Vendedor", "FUNC-001")
	f2 := models.NewFuncionario("Mariana Lima", "mariana.l", "123", "gerente", "11966667777", "mariana.l@comercialsilva.com", "Gerente de Vendas", "FUNC-002")
	SalvarFuncionario(f1)
	SalvarFuncionario(f2)
}
