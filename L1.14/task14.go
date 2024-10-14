package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

// Определение типа с использованием утверждений типов
func printTypeWithAssertion(value interface{}) {
	op := "Using Type Assertion:"

	// При таком подходе для каждого конкретного типа канала необходимо реализовать обработку, также для (chan<- и <-chan)
	switch v := value.(type) {
	case int:

		fmt.Printf("%s\nType: %s\nValue: %d\n", op, "int", v)
	case string:
		fmt.Printf("%s\nType: %s\nValue: \"%s\"\n", op, "string", v)
	case bool:
		fmt.Printf("%s\nType: %s\nValue: %t\n", op, "bool", v)
	case chan int:
		fmt.Printf("%s\nType: %s\n", op, "chan int")
	case chan string:
		fmt.Printf("%s\nType: %s\n", op, "chan string")
	case chan bool:
		fmt.Printf("%s\nType: %s\n", op, "chan bool")
	case chan interface{}:
		fmt.Printf("%s\nType: %s\n", op, "chan interface{}")
	case chan struct{}:
		fmt.Printf("%s\nType: %s\n", op, "chan struct{}")
	default:
		fmt.Printf("%s\nUnknown type: %T\n", op, v)
	}
	fmt.Println("--------------------------")
}

// Определение типа с использованием рефлексии
func printTypeWithReflection(value interface{}) {
	op := "Using Reflection:"
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	if t.Kind() == reflect.Chan {
		fmt.Printf("%s\nType: %s\n", op, t)
	} else {
		fmt.Printf("%s\nType: %s\nValue: %v\n", op, t, v)
	}
	fmt.Println("--------------------------")
}

func main() {
	testValues := []interface{}{
		17,
		"ol",
		false,
		make(chan int),
		make(chan string),
		make(chan interface{}),
	}

	for _, value := range testValues {
		printTypeWithAssertion(value)
		printTypeWithReflection(value)
	}
}
