package main

import "fmt"

type Product interface {
	doStuff() string
}
type Items struct {
	name string
}

func (b *Items) doStuff() string {
	return "this is " + b.name
}

type ItemsA struct {
	Items
}

func NewProductA() Product {
	return &ItemsA{
		Items: Items{
			name: "Items A",
		},
	}
}

type ItemsB struct {
	Items
}

func NewProductB() Product {
	return &ItemsB{
		Items: Items{
			name: "Items B",
		},
	}
}

func CreateProduct(name string) Product {
	if name == "A" {
		return NewProductA()
	} else if name == "B" {
		return NewProductB()
	}
	return nil
}
func main() {
	productA := CreateProduct("A")
	fmt.Println(productA.doStuff())
	productB := CreateProduct("B")
	fmt.Println(productB.doStuff())
}
