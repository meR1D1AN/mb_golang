package main

import (
	"fmt"
	"github.com/google/uuid"
)

type Person struct {
	id      uuid.UUID
	name    string
	surname string
	age     int
	email   string
	phone   string
}

func (p Person) greet() {
	fmt.Printf("Привет, меня зовут %v!\n", p.name)
}

func main() {
	p := Person{
		id:      uuid.New(),
		name:    "John",
		surname: "Doe",
		age:     42,
		email:   "john.doe@example.com",
		phone:   "+79958878899",
	}
	person1 := Person{
		id:      uuid.New(),
		name:    "John",
		surname: "Doe",
		age:     42,
		email:   "john.doe@example.com",
		phone:   "+79958878899",
	}
	fmt.Println(p)
	fmt.Println(person1.age)
	person1.age = 25
	fmt.Println(person1.age)
	person1.greet()
}
