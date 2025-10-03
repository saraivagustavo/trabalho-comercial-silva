package classes

type Funcionario struct {
	Usuario
	cargo     string
	matricula string
}

// ===== CONSTRUTOR DA CLASSE =====
func NewFuncionario(nome, login, senhaPura, tipo, telefone, email, cargo, matricula string) *Funcionario {
	funcionario := &Funcionario{}

	funcionario.SetNome(nome)
	funcionario.SetLogin(login)
	funcionario.SetTipo(tipo)
	funcionario.SetTelefone(telefone)
	funcionario.SetEmail(email)
	funcionario.SetCargo(cargo)
	funcionario.SetMatricula(matricula)

	return funcionario
}

// ===== GETTERS para os campos =====
func (funcionario *Funcionario) GetCargo() string {
	return funcionario.cargo
}

func (funcionario *Funcionario) GetMatricula() string {
	return funcionario.matricula
}

// ===== SETTERS para atribuir valor pros campos =====
func (funcionario *Funcionario) SetCargo(cargo string) {
	funcionario.cargo = cargo
}

func (funcionario *Funcionario) SetMatricula(matricula string) {
	funcionario.matricula = matricula
}

func (f *Funcionario) ToString() string {
	return "--- DADOS DO FUNCIONÁRIO ---\n" +
		"Nome: " + f.GetNome() + "\n" +
		"Matrícula: " + f.GetMatricula() + "\n" +
		"Cargo: " + f.GetCargo() + "\n" +
		"Telefone: " + f.TelefoneFormatado() + "\n" +
		"Email: " + f.GetEmail() + "\n" +
		"----------------------------"
}
