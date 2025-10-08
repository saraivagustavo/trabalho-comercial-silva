package controllers

import (
	"strings"
	"trabalho/database"
	"trabalho/models"
	"trabalho/views"
)

// mesma lógica do cliente_controller
func CadastrarFuncionario() {
	// 1. view coleta os dados
	nome, login, tipo, telefone, email, cargo, matricula := views.SolicitarDadosFuncionario()

	//tem que ver como vai ser a senha depois
	senhaTeste := "123456"

	//limpar as strings também
	nome = strings.TrimSpace(nome)
	login = strings.TrimSpace(login)
	tipo = strings.TrimSpace(tipo)
	telefone = strings.TrimSpace(telefone)
	email = strings.TrimSpace(email)
	cargo = strings.TrimSpace(cargo)
	matricula = strings.TrimSpace(matricula)

	// 2. model cria o objeto
	novoFuncionario := models.NewFuncionario(nome, login, senhaTeste, tipo, telefone, email, cargo, matricula)

	// 3. database salva
	database.SalvarFuncionario(novoFuncionario)

	// 4. view exibe o resultado do cadastro
	views.ExibirFuncionario(novoFuncionario)
}

func ListarFuncionarios() {
	funcionarios := database.FuncionariosDB
	views.ExibirListaFuncionarios(funcionarios)
}
