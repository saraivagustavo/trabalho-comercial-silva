package models

type Funcionario struct {
	Usuario
	cargo     string
	matricula string
}

// ===== CONSTRUTOR DA CLASSE =====
func NewFuncionario(nome, login, senha, tipo, telefone, email, cargo, matricula string) *Funcionario {
	return &Funcionario{
		Usuario: Usuario{
			nome:     nome,
			login:    login,
			senha:    senha,
			tipo:     tipo,
			telefone: telefone,
			email:    email,
		},
		cargo:     cargo,
		matricula: matricula,
	}
}

// ===== GETTERS para os campos =====
func (funcionario *Funcionario) GetCargo() string {
	return funcionario.cargo
}

func (funcionario *Funcionario) GetMatricula() string {
	return funcionario.matricula
}

// ===== Setters da classe =====
func (funcionario *Funcionario) SetCargo(novoCargo string) {
	funcionario.cargo = novoCargo
}

// formatação dos dados do funcionário
func (f *Funcionario) ToString() string {
	return "\n--- DADOS DO FUNCIONÁRIO ---\n" +
		"Nome: " + f.GetNome() + "\n" +
		"Matrícula: " + f.GetMatricula() + "\n" +
		"Cargo: " + f.GetCargo() + "\n" +
		"Telefone: " + f.TelefoneFormatado() + "\n" +
		"Email: " + f.GetEmail() + "\n" +
		"----------------------------"
}
