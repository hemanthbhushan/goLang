package main

import "fmt"

func calculation(a, b int) {
	cal := a + b

	fmt.Println(cal)

}

func check(y []int, x func(a, b int)) string {

	for index, value := range y {
		x(index, value)

	}
	return "complete"
}

func main() {

	arr := []int{1, 2, 3, 4}

	fmt.Println(check(arr, calculation))

}
