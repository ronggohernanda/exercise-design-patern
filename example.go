package main

import "fmt"

type User interface {
	doStuff() string
}

type Itsme struct {
	name string
}

func (i *Itsme) doStuff() string {
	return "hello there, i am " + i.name
}

type Admin struct {
	Itsme
}

func NewAdmin() User {
	return &Admin{
		Itsme: Itsme{
			name: "Admin",
		},
	}
}

type Customer struct {
	Itsme
}

func NewCustomer() User {
	return &Customer{
		Itsme: Itsme{
			name: "Customer",
		},
	}
}

func CreateUser(name string) User {
	if name == "Admin" {
		return NewAdmin()
	} else if name == "Customer" {
		return NewCustomer()
	}
	return nil
}

func main() {
	admin := CreateUser("Admin")
	fmt.Println(admin.doStuff())
	customer := CreateUser("Customer")
	fmt.Println(customer.doStuff())
}
