package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}

	a := person{name: "lee", age: 28}
	fmt.Println(s.name)

	var spp *person = &a

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

	fmt.Println("person a :", a.name)
	fmt.Println("person a :", spp.age)
	fmt.Println("person a :", spp.age)
}
