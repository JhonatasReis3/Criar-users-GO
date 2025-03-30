package main

import "fmt"

func main() {
	apresentacao()
}
func apresentacao() {
	fmt.Println("---Bem vindo ao seu sistema de cadastro simples---")
	fmt.Println("\n" + "Por Gentileza escolha uma opção:")
	var EscolhaMenuOpicional int
	OpcaoMenuInicial := [...]string{"1-Login Gerente", "2-Login usuario", "3-Cadastro novo usuario", "4- Sair"}
	for i := 0; i < len(OpcaoMenuInicial); i++ {
		fmt.Println(OpcaoMenuInicial[i])
	}
	fmt.Scanln(&EscolhaMenuOpicional)
	fmt.Println("Sua escolha foi:", EscolhaMenuOpicional)
}
