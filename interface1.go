package main

import "fmt"

type people interface {
	greeting()
	calling() string
}

type man struct {
	name   string
	gender string
	age    int64
}

type woman struct {
	name   string
	gender string
	age    int64
}

func (r man) greeting() {
	fmt.Println("hello I'm Mr.", r.name)
}

func (r woman) greeting() {
	fmt.Println("hello I'm Ms.", r.name)
}

func (r man) calling() string {
	if r.age > 20 {
		return "gentle man"
	} else {
		return "boy"
	}
}

func (r woman) calling() string {
	if r.age > 20 {
		return "lady"
	} else {
		return "girl"
	}
}

func person_info(info people) {
	info.greeting()
	fmt.Println(info.calling())
}

func main() {
	leeyunhyuk := man{name: "leeyunhyuk", gender: "male", age: 28}
	parkheejun := woman{name: "parkheejun", gender: "female", age: 29}
	namgeupsik := man{name: "namgeupsik", gender: "male", age: 18}
	yeojammin := woman{name: "yeojammin", gender: "female", age: 13}

	person_info(leeyunhyuk)
	person_info(parkheejun)
	person_info(namgeupsik)
	person_info(yeojammin)
}
