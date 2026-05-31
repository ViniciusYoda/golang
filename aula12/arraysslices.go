package main

import "fmt"

func main(){
	var array [2] string
	array[0] = "Go"
	array[1] = "Lang"
	fmt.Printf("O array é: %v\n", array)
	slice := []string{"Go", "Lang"}
	fmt.Printf("O slice é: %v\n", slice)
	
}