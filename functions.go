// package main

// import "fmt"

// func check() int {
// 	return 2
// }

// func calculate(x, y, i int) {
// 	calc := x + y
// 	fmt.Printf("total value is = %d with index %d \n ", calc, i)

// }

// func mathOperations(x []int, cal func(x, y, i int)) string {
// 	for index, value := range x {
// 		cal(value, value, index)
// 	}

// 	return "completed"
// }

// func main() {

// 	x := check()
// 	fmt.Println(x)
// 	// fmt.Println(calculate(2, 3))

// 	fmt.Println(mathOperations([]int{1, 2, 3, 4}, calculate))

// }
