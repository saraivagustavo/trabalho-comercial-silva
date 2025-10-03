package classes

import "fmt"

type Usuario struct {
	nome     string
	login    string
	senha    string
	tipo     string
	telefone string
	email    string
}

// ===== GETTERS para os campos =====
func (usuario *Usuario) GetNome() string {
	return usuario.nome
}

func (usuario *Usuario) GetLogin() string {
	return usuario.login
}

func (usuario *Usuario) GetTipo() string {
	return usuario.tipo
}

func (usuario *Usuario) GetTelefone() string {
	return usuario.telefone
}

func (usuario *Usuario) GetEmail() string {
	return usuario.email
}

// ===== SETTERS para atribuir valor pros campos =====
func (usuario *Usuario) SetNome(nome string) {
	usuario.nome = nome
}

func (usuario *Usuario) SetLogin(login string) {
	usuario.login = login
}

func (usuario *Usuario) SetTipo(tipo string) {
	usuario.tipo = tipo
}

func (usuario *Usuario) SetTelefone(telefone string) {
	usuario.telefone = telefone
}

func (usuario *Usuario) SetEmail(email string) {
	usuario.email = email
}

// ===== Métodos da classe =====
func (u *Usuario) TelefoneFormatado() string {
	return fmt.Sprintf("(%s) %s-%s", u.telefone[0:2], u.telefone[2:7], u.telefone[7:11])
}
