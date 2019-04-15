package main

import "fmt"

// var x int

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning. Person."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Good morning. Agent."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// x = 7
	// Print var type
	// fmt.Printf("%T", x)

	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"Ivan": 26,
		"Job":  42,
	}
	fmt.Println(m)

	p1 := person{
		"Ivan",
		"Satyaputra",
	}
	p1.speak()

	sa1 := secretAgent{
		person{"James", "Bond"},
		true,
	}
	sa1.speak()
	sa1.person.speak()

	saySomething(p1)
	saySomething(sa1)
}
