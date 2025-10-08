package views

import "fmt"

// menu principal, dividir em submenus para não ficar poluído
func MostrarMenuPrincipal() {
	fmt.Println("\n====== COMERCIAL SILVA | MENU PRINCIPAL ======")
	fmt.Println("1. Gerenciar Clientes")
	fmt.Println("2. Gerenciar Funcionários")
	fmt.Println("3. Gerenciar Produtos")
	fmt.Println("4. Gerenciar Pedidos")
	fmt.Println("0. Sair")
	fmt.Println("============================================")
	fmt.Print("Opção: ")
}

// submenu genérico, recebe o nome da entidade que vai gerenciar como parâmetro
func MostrarSubMenu(entidade string) {
	fmt.Printf("\n--- Gerenciar %s ---\n", entidade)
	fmt.Printf("1. Adicionar %s", entidade)
	fmt.Printf("\n2. Listar %s", entidade)
	fmt.Printf("\n3. Editar %s", entidade)
	fmt.Printf("\n4. Deletar %s", entidade)
	fmt.Println("\n0. Voltar ao Menu Principal")
	fmt.Println("---------------------------")
	fmt.Print("Opção: ")
}
