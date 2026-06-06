package main

import "fmt"

func main() {
	test := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}

	fmt.Printf("%T", test)

	var test2 interface{}

	test2 = 20.5

	testJson := map[string]interface{}{
		"nome":      "João",
		"sobrenome": "Silva",
		"idade":     30,
	}

	fmt.Printf("%T", test2)
	fmt.Printf("%T", testJson)

	var test3 [4]string = [4]string{"a", "b", "c", "d"}

	fmt.Printf("%T", test3)

	fmt.Println(len(test3))
	fmt.Println(cap(test3))

	var u user = user{
		name:    "João",
		surname: "Silva",
		age:     30,
	}

	fmt.Printf("%T", u)
	fmt.Printf("%+v", u)
	fmt.Printf("%v", u)
	fmt.Printf("%#v", u)

}

type User struct {
	Email    string
	password string
}

type user struct {
	name    string
	surname string
	age     int
}
