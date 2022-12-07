package main

import "fmt"

func main() {

	//arrays
	var arr1 [4]int = [4]int{1, 2, 3, 4}

	//slices
	//append works only on the slices just like a push for arr
	var arr []int = []int{1, 2, 3, 4, 5}
	arr = append(arr, 23)

	//slice

	rangeOne := arr1[1:3]
	rangeTwo := arr1[:4]
	rangeThree := arr[1:4]
	//range one [2 3]
	//range two [1 2 3 4]
	//range three [1 2 3 4]

	fmt.Println("range one", rangeOne)
	fmt.Println("range two", rangeTwo)
	fmt.Println("range three", rangeThree)

	fmt.Println("arr,arr1", arr, arr1)
}
