package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 1. ESTRUTURA DOS DADOS (Equivalente à tabela do banco de dados)
type Registro struct {
	Nome   string
	CEP    string
	Estado string
	Pais   string
}

// 2. LEITURA E VALIDAÇÃO (Passagem eficiente de ponteiro para o scanner)
func lerCampoValido(scanner *bufio.Scanner, rotulo string, minTam int, maxTam int) string {
	fmt.Printf("%s: ", rotulo)
	scanner.Scan()
	texto := strings.TrimSpace(scanner.Text())

	for len(texto) < minTam || len(texto) > maxTam {
		fmt.Printf("Entrada inválida (%s deve ter entre %d e %d caracteres).\n", rotulo, minTam, maxTam)
		fmt.Printf("Digite novamente %s: ", rotulo)
		scanner.Scan()
		texto = strings.TrimSpace(scanner.Text())
	}

	return texto
}

// 3. CARREGAR DADOS DO DISCO PARA A MEMÓRIA (Ao Iniciar)
func carregarDeCSV() []Registro {
	var registros []Registro

	arquivo, erro := os.Open("cadastros.csv")
	if erro != nil {
		// Se o arquivo não existir (primeira execução), retorna a lista vazia
		return registros
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	primeiraLinha := true

	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())

		if len(linha) == 0 {
			continue
		}

		// Pula o cabeçalho ("Nome;CEP;Estado;Pais")
		if primeiraLinha {
			primeiraLinha = false
			continue
		}

		// Divide os campos pelo ponto e vírgula ';'
		campos := strings.Split(linha, ";")

		if len(campos) == 4 {
			registroLido := Registro{
				Nome:   campos[0],
				CEP:    campos[1],
				Estado: campos[2],
				Pais:   campos[3],
			}
			registros = append(registros, registroLido)
		}
	}

	return registros
}

// 4. SALVAR A MEMÓRIA NO DISCO (Arquivo CSV)
func salvarEmCSV(registros []Registro) {
	arquivo, erro := os.Create("cadastros.csv")
	if erro != nil {
		fmt.Println("Erro ao salvar o arquivo CSV:", erro)
		return
	}
	defer arquivo.Close()

	// Escreve o cabeçalho para planilhas (Excel / LibreOffice)
	fmt.Fprintln(arquivo, "Nome;CEP;Estado;Pais")

	for _, reg := range registros {
		linha := fmt.Sprintf("%s;%s;%s;%s", reg.Nome, reg.CEP, reg.Estado, reg.Pais)
		fmt.Fprintln(arquivo, linha)
	}

	fmt.Println("Dados salvos com sucesso em 'cadastros.csv'!")
}

// 5. REMOVER REGISTRO DA MEMÓRIA
func removerRegistro(scanner *bufio.Scanner, registros []Registro) []Registro {
	if len(registros) == 0 {
		fmt.Println("Nenhum registro disponível para remoção.")
		return registros
	}

	fmt.Println("\n--- SELECIONE O REGISTRO PARA REMOVER ---")
	for i, reg := range registros {
		// i+1 apenas para apresentação humana na tela
		fmt.Printf("[%d] Nome: %s | CEP: %s\n", i+1, reg.Nome, reg.CEP)
	}

	fmt.Print("Digite o número do registro que deseja apagar (ou 0 para cancelar): ")
	scanner.Scan()
	entrada := strings.TrimSpace(scanner.Text())

	opcaoNumero := 0
	fmt.Sscanf(entrada, "%d", &opcaoNumero)

	if opcaoNumero == 0 {
		fmt.Println("Operação cancelada.")
		return registros
	}

	// Converte do número na tela (1, 2, 3...) para o índice da memória (0, 1, 2...)
	indiceMemoria := opcaoNumero - 1

	// Proteção contra erro de índice fora de alcance (panic/out of range)
	if indiceMemoria < 0 || indiceMemoria >= len(registros) {
		fmt.Println("Número inválido! Nenhum registro foi removido.")
		return registros
	}

	nomeRemovido := registros[indiceMemoria].Nome

	// Fatiamento de slice para remover a posição desejada
	registros = append(registros[:indiceMemoria], registros[indiceMemoria+1:]...)

	fmt.Printf("Registro de '%s' removido da memória!\n", nomeRemovido)

	// Salva automaticamente o CSV atualizado após deletar
	salvarEmCSV(registros)

	return registros
}

// 6. FUNÇÃO PRINCIPAL (Controle de Fluxo)
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Carrega dados previamente salvos
	bancoDeDados := carregarDeCSV()

	if len(bancoDeDados) > 0 {
		fmt.Printf(" %d registro(s) recuperado(s) do arquivo 'cadastros.csv'.\n", len(bancoDeDados))
	}

	opcao := ""

	for opcao != "0" {
		fmt.Println("\n===================================")
		fmt.Println("     SISTEMA DE CADASTRO (CEP)     ")
		fmt.Println("===================================")
		fmt.Println("1. Cadastrar nova pessoa")
		fmt.Println("2. Listar cadastros na tela")
		fmt.Println("3. Salvar/Exportar para arquivo CSV")
		fmt.Println("4. Remover um cadastro")
		fmt.Println("0. Sair do programa")
		fmt.Print("Escolha uma opção: ")

		scanner.Scan()
		opcao = strings.TrimSpace(scanner.Text())

		switch opcao {
		case "1":
			fmt.Println("\n--- NOVO CADASTRO ---")
			nome := lerCampoValido(scanner, "Nome", 3, 50)
			cep := lerCampoValido(scanner, "CEP (8 números)", 8, 8)
			estado := lerCampoValido(scanner, "Estado (ex: RS, SP)", 2, 2)
			pais := lerCampoValido(scanner, "País", 3, 30)

			novoRegistro := Registro{
				Nome:   nome,
				CEP:    cep,
				Estado: estado,
				Pais:   pais,
			}

			bancoDeDados = append(bancoDeDados, novoRegistro)
			fmt.Println("Registro adicionado à memória!")

		case "2":
			fmt.Println("\n--- REGISTROS NA MEMÓRIA ---")
			if len(bancoDeDados) == 0 {
				fmt.Println("Nenhum registro encontrado na memória.")
			} else {
				for i, reg := range bancoDeDados {
					fmt.Printf("%d. Nome: %s | CEP: %s | UF: %s | País: %s\n",
						i+1, reg.Nome, reg.CEP, reg.Estado, reg.Pais)
				}
			}

		case "3":
			if len(bancoDeDados) == 0 {
				fmt.Println("Nenhum dado na memória para exportar.")
			} else {
				salvarEmCSV(bancoDeDados)
			}

		case "4":
			bancoDeDados = removerRegistro(scanner, bancoDeDados)

		case "0":
			fmt.Println("\nEncerrando o sistema. Até logo!")

		default:
			fmt.Println("Opção inválida! Escolha uma opção de 0 a 4.")
		}
	}
}
