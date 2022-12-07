package main

import "fmt"

func makesum(a int, b int) (int, string) {
	if (a + b) == 1 {
		return a + b, "number 1"
	}

	if (a + b) != 1 {
		return a + b, "this is not 1"
	}

	return 100, "test"
}

func callback(test int, f func(*int)) {
	fmt.Println("call back & number : ", test)
	f(&test)
	fmt.Println("after call back & number : ", test)
}

func mul_reference(num *int) {
	*num = *num * 10
}

func main() {

	a, b := makesum(1, 1)
	_, c := makesum(1, 2)

	fmt.Println(a, b)
	fmt.Println(c)

	val := 5
	var ptr *int = &val
	callback(*ptr, mul_reference)

}
