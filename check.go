package main

import "fmt"

func main() {

	//arrays
	var arr1 [4]int = [4]int{1, 2, 3, 4}

	//slices
	//append works only on the slices just like a push for arr
	var arr []int = []int{1, 2, 3, 4, 5}
	arr = append(arr, 23)

	fmt.Println("arr,arr1", arr, arr1)
}
