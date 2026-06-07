package main

import "fmt"

func main(){

	testParametros2 := func(){
		fmt.Println("test2")
	}
	testParametro := func() {
		fmt.Println("test")
	}

	test(testParametro, testParametros2)
}

func test(valoresString ...func()) {
	for _, x := range valoresString {
		x()
	}
}