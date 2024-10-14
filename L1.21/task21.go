package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

/*
Реализовать паттерн «адаптер» на любом примере.
*/

type Person struct {
	Name string
	Age  int
}

// Интерфейс, который адаптирует разные форматы вывода
type Formatter interface {
	Format(Person) string
}

// Реализация JSON формата
type JSONFormatter struct{}

func (j *JSONFormatter) Format(p Person) string {
	data, _ := json.Marshal(p)
	return string(data)
}

// Реализация XML формата
type XMLFormatter struct{}

func (x *XMLFormatter) Format(p Person) string {
	data, _ := xml.MarshalIndent(p, "", "  ") // Используем MarshalIndent для отступов
	return string(data)
}
func printFormatted(f Formatter, p Person) {
	fmt.Println(f.Format(p))
}

func main() {
	p := Person{Name: "John", Age: 30}

	fmt.Println("JSON Formatter")
	jsonFormatter := &JSONFormatter{}
	printFormatted(jsonFormatter, p)

	fmt.Println("XML Formatter")
	xmlFormatter := &XMLFormatter{}
	printFormatted(xmlFormatter, p)
}
