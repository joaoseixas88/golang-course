package main

import "fmt"

type Person struct {
	name    string
	surname string
	age     int16
}

type Student struct {
	Person
	course     string
	university string
}

func main() {

	person := Student{
		Person: Person{
			name:    "Joao",
			surname: "Seixas",
			age:     35,
		},
		course:     "Developer",
		university: "Udemy",
	}

	fmt.Println(person.name)
	fmt.Println(person.course)

}
