package main

import "fmt"

func main() {

	var variavel int = 10
	var variavel2 int = variavel

	var memory1 *int = &variavel
	var memory2 *int = &variavel2

	variavel++

	fmt.Println(variavel, variavel2)
	fmt.Println(memory1, memory2)

}
