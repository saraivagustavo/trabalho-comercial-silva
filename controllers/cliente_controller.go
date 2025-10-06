package controllers

import (
	"trabalho/database"
	"trabalho/models"
	"trabalho/views"
	"strings"
)

//a função CadastrarCliente vai gerenciar o fluxo de cadastro de um cliente
//ela orquestra a interação entre a View e o Model
func CadastrarCliente() {
	//1. usa a função de dentro da viiew pra solicitar os dados do cliente
	nome, login, tipo, telefone, email, identificacao, endereco := views.SolicitarDadosCliente()

	//tava dando erro nas formatações pq as strings vinham com o '\n' do enter, tem que "limpar" elas
	nome = strings.TrimSpace(nome)
	login = strings.TrimSpace(login)
	tipo = strings.TrimSpace(tipo)
	telefone = strings.TrimSpace(telefone)
	email = strings.TrimSpace(email)
	identificacao = strings.TrimSpace(identificacao)
	endereco = strings.TrimSpace(endereco)

	//2. depois de pegar os dados, manda pro model criar um novo cliente usando o construtor recebendo os parâmetros
	novoCliente := models.NewCliente(nome, login, tipo, telefone, email, identificacao, endereco)

	//3. salva o cliente no banco
	database.SalvarCliente(novoCliente)

	//4. printa a confirmação de que cadastrou o cliente
	views.ExibirCliente(novoCliente)
}
