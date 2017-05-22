package main

import "fmt"

/**
 * RESULT

arrLength :  5
plusNum :  10

 */
func main() {
	plus := func(numbers ... int) (int, int) {
		rtnInt := 0;
		for _, number := range numbers {
			rtnInt += number
		}

		return len(numbers), rtnInt

	}

	arrLength, plusNum := plus(0, 1, 2, 3, 4);

	fmt.Println("arrLength : ", arrLength)
	fmt.Println("plusNum : ", plusNum)
}
