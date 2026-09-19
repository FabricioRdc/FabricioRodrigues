package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func lerNomeCompleto(scanner *bufio.Scanner) string {
	fmt.Println("Digite seu nome completo:")
	scanner.Scan()
	nome := strings.TrimSpace(scanner.Text())

	// Validações:
	// 1. Tamanho mínimo e máximo
	// 2. Presença de pelo menos um espaço (para garantir nome + sobrenome)
	for len(nome) < 3 || len(nome) > 100 || !strings.Contains(nome, " ") {
		fmt.Println("\nNome inválido!")
		fmt.Println("- Deve ter entre 3 e 100 caracteres.")
		fmt.Println("- Deve conter pelo menos um espaço (ex: 'Maria Silva').")
		fmt.Println("Por favor, digite novamente:")
		
		scanner.Scan()
		nome = strings.TrimSpace(scanner.Text())
	}

	return nome
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	nomeCompleto := lerNomeCompleto(scanner)

	fmt.Println("\nCadastro realizado!")
	fmt.Println("Nome completo registrado:", nomeCompleto)
}
