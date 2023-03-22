package main

import "fmt"

func main() {
	array := make([]int, 10, 15)
	array[1] = 2
	array[9] = 2
	// array2 := []int{2,3}
	// append(array, ...array2)
	array = append(array, 2)
	array = append(array, 2)
	array = append(array, 2)
	array = append(array, 2)
	array = append(array, 2)
	array = append(array, 2)

	fmt.Println(array)
	fmt.Println(cap(array))
}
