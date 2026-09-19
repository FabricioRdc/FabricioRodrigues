package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Qual é o seu nome?")
	scanner.Scan()
	nome := strings.TrimSpace(scanner.Text())
	fmt.Println("Qual é o seu sobrenome?")
	scanner.Scan()
	sobrenome := strings.TrimSpace(scanner.Text())
	fmt.Println("O seu nome completo é:", nome, sobrenome)
}
