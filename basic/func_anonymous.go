package main

import "fmt"

func main() {
	plus := func(numbers ... int) (int, int) {
		rtnInt := 0;
		for _, number := range numbers {
			rtnInt += number
		}

		return len(numbers), rtnInt

	}

	arrLength, plusNum := plus(0, 1, 2, 3, 4);

	fmt.Println("arrLenth : ", arrLength)
	fmt.Println("plusNum : ", plusNum)
}
