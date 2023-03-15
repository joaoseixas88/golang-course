package main

import "fmt"

func main() {

	var variavel string = "Something"
	var variable int
	var (
		variavel1 string
		variavel4 int
	)
	variable = 32
	variavel2 := "Variavel 2"

	const constante1 string = "Constante 1"
	const constante2 = "Constante2"

	fmt.Println(variavel)
	fmt.Println(variable)
	fmt.Println(variavel2)
	fmt.Println(variavel1)
	fmt.Println(variavel4)
	fmt.Println(constante1)
	fmt.Println(constante2)
}
