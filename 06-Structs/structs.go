package main

import "fmt"

// Structs sao o mais proximo de classes que o go tem

type Person struct {
	name    string
	surname string
	age     int
	job     Job
}

type Job struct {
	name          string
	averageSalary float64
}

func main() {

	joao := Person{
		name:    "Joao",
		surname: "Seixas",
		age:     35,
		job:     Job{name: "Developer", averageSalary: 3500.00},
	}

	fmt.Println(joao)
	fmt.Printf("%T", joao)
}
