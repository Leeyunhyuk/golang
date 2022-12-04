package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r rect) value_() int {
	r.width = r.width * 2
	return r.width
}

func (r *rect) pointer_() int {
	r.width = r.width * 2
	return r.width
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("value :", r.value_())
	fmt.Println("value value:", r.width)

	fmt.Println("address :", r.pointer_())
	fmt.Println("address value:", r.width)
}
