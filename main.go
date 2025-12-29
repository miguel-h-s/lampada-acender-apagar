package main

import "fmt"

var ligada bool

func cabecalho() {
	fmt.Println("=== ACENDA/APAGUE A LÂMPADA ===")
}

func opcoes() {
	fmt.Println("1. ACENDER LÂMPADA")
	fmt.Println("2. APAGAR LÂMPADA")
	fmt.Println("0. SAIR")
	fmt.Printf("estado: %t\n", ligada)
}

func ligar() {
	fmt.Println("voce acendeu a lâmpada!")
	fmt.Printf("estado: %t\n", ligada)
}

func desligar() {
	fmt.Println("você desligou a lâmpada!")
	fmt.Printf("estado: %t\n", ligada)
}

func separador() {
	fmt.Println("===============================")
}

func main() {
	ligada = false

	cabecalho()
	opcoes()
	separador()

	var opcao int

	for {
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			ligar()
			separador()
		case 2:
			desligar()
			separador()
		case 0:
			separador()
			return
		}

	}
}
