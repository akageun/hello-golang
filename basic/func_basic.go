package main

import "fmt"

func plus(val1 int, val2 int) int {
	return val1 + val2
}

func subtract(val1 int, val2 int) int {
	return val1 - val2
}

/**
 * RESULT

plus :  15
subtract :  5

 */
func main() {

	a := 10
	b := 5

	fmt.Println("plus : ", plus(a, b))
	fmt.Println("subtract : ", subtract(a, b))
}
