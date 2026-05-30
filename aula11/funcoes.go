package main

import "fmt"

func main() {
	resultado := soma(3, 5)
	fmt.Printf("A soma de 3 e 5 é: %d\n", resultado)
	resultado = subtracao(10, 4)
	fmt.Printf("A subtração de 10 e 4 é: %d\n", resultado)
	resultado = multiplicacao(6, 7)
	fmt.Printf("A multiplicação de 6 e 7 é: %d\n", resultado)
	resultado, err := divisao(20, 4)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("A divisão de 20 por 4 é: %d\n", resultado)
	}
	resultado, err = divisao(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("A divisão de 10 por 0 é: %d\n", resultado)
	}
	printaNome("GoLang")
	nome, sobrenome := "Go", "Lang"
	fmt.Printf("O nome completo é: %s %s\n", nome, sobrenome)
}

func soma(x int, y int) int {
	return x + y
}	

func subtracao(x int, y int) int {
	return x - y
}

func multiplicacao(x int, y int) int {
	return x * y
}

func divisao(x int, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("não é possível dividir por zero")
	}	
	return x / y, nil
}

func printaNome(nome string) {
	fmt.Printf("O nome é: %s\n", nome)
}

func printaNomeCompleto(nome string, sobrenome string) {
	fmt.Printf("O nome completo é: %s %s\n", nome, sobrenome)
}