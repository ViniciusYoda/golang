package main

import "fmt"

func main() {
	idade := map[string]int{
		"João": 30,
		"Maria": 25,
	}
	fmt.Printf("A idade de João é: %d\n", idade["João"])
	fmt.Printf("A idade de Maria é: %d\n", idade["Maria"])

	anoNasc := map[string]int{
		"João": 1990,
		"Maria": 1995,
	}
	fmt.Printf("O ano de nascimento de João é: %d\n", anoNasc["João"])
	fmt.Printf("O ano de nascimento de Maria é: %d\n", anoNasc["Maria"])
	anoNasc["Pedro"] = 2000
	fmt.Printf("O ano de nascimento de Pedro é: %d\n", anoNasc["Pedro"])


}