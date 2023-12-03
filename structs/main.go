package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	//Zero Value initialization
	//var alex person

	alex := person{firstName: "Alex", lastName: "Anderson"}
	alex.firstName = "Alexis"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
