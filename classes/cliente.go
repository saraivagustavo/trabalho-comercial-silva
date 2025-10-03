package classes

import "fmt"

type Cliente struct {
	Usuario
	identificacao string
	endereco      string
}

// ===== CONSTRUTOR DA CLASSE =====
func NewCliente(nome, login, tipo, telefone, email, identificacao, endereco string) *Cliente {
	cliente := &Cliente{}
	
	cliente.SetNome(nome)
	cliente.SetLogin(login)
	cliente.SetTipo(tipo)
	cliente.SetTelefone(telefone)
	cliente.SetEmail(email)
	cliente.SetIdentificacao(identificacao)
	cliente.SetEndereco(endereco)

	return cliente
}

// ===== GETTERS para os campos =====
func (cliente *Cliente) GetIdentificacao() string {
	return cliente.nome
}

func (cliente *Cliente) GetEndereco() string {
	return cliente.endereco
}

// ===== SETTERS para atribuir valor pros campos =====
func (cliente *Cliente) SetIdentificacao(identificacao string) {
	cliente.identificacao = identificacao
}

func (cliente *Cliente) SetEndereco(endereco string) {
	cliente.endereco = endereco
}

// ===== Métodos da classe =====
func (c *Cliente) IdentificacaoFormatada() string {
	id := c.GetIdentificacao()
	//formatação pra CPF
	if len(id) == 11 {
		return fmt.Sprintf("%s.%s.%s-%s", id[0:3], id[3:6], id[6:9], id[9:11])
	}
	//formatação pra CNPJ
	if len(id) == 14 {
		return fmt.Sprintf("%s.%s.%s/%s-%s", id[0:2], id[2:5], id[5:8], id[8:12], id[12:14])
	}
	return id
}

func (c *Cliente) ToString() string {
	return "--- DADOS DO CLIENTE ---\n" +
		"Nome: " + c.GetNome() + "\n" +
		"Documento: " + c.IdentificacaoFormatada() + "\n" +
		"Telefone: " + c.TelefoneFormatado() + "\n" +
		"Email: " + c.GetEmail() + "\n" +
		"Endereço: " + c.GetEndereco() + "\n" +
		"-------------------------"
}
