package main

import "fmt"


/**
 * RESULT

arr 5 :  8
length :  10
arr2 :  0
arr2 :  1
arr2 :  2
arr2 :  3
arr2 :  4
arr2 End

 */
func main() {
	var arr [10]int

	arr[5] = 8
	fmt.Println("arr 5 : ", arr[5])
	fmt.Println("length : ", len(arr))

	arr2 := [5]int{0, 1, 2, 3, 4}
	for i := 0; i < len(arr2); i++ {
		fmt.Println("arr2 : ", arr2[i])

		if arr2[i] == 4 {
			fmt.Println("arr2 End")
		}
	}

}
