package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human
(аналог наследования).
*/

type Human struct {
	name string
	age  int
}

func (human Human) Greeting() string {
	return fmt.Sprintf("Hi, I'm %s! I am %d years old.", human.name, human.age)
}

type Action struct {
	Human       // Анонимное встраивание структуры Human
	alias Human // Именованное встраивание структуры Human
}

func main() {
	a := Action{
		Human: Human{name: "Kakashi", age: 26},
		alias: Human{name: "Sukea", age: 26},
	}

	// Пример использования анонимного встраивания
	fmt.Println(a.Greeting())

	// Пример использования именованного встраивания
	fmt.Println(a.alias.Greeting())
}
