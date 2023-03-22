package main

import "fmt"

func main() {

	position := 45

	for i := 1; i <= position; i++ {
		fmt.Println(fibonacci(i))
	}

}

func fibonacci(position int) int {
	if position <= 1 {
		return 1
	}
	return fibonacci(position-2) + fibonacci(position-1)

}
