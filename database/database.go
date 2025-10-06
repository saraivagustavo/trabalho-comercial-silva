package database

import "trabalho/models"

// ClientesDB é o banco pra armazenar os clientes
// A chave do mapa vai ser o cpf ou cnpj do cliente
var ClientesDB = make(map[string]*models.Cliente) //usar map aqui pq é mais fácil de buscar depois

//mesma coisa, mas pros funcionários e com a chaave sendo a matrícula dele
var FuncionariosDB = make(map[string]*models.Funcionario)

//a função SalvarCliente adiciona um novo cliente no banco
func SalvarCliente(cliente *models.Cliente) {
	ClientesDB[cliente.GetIdentificacao()] = cliente
}

//" " " "
func SalvarFuncionario(funcionario *models.Funcionario) {
	FuncionariosDB[funcionario.GetMatricula()] = funcionario
}