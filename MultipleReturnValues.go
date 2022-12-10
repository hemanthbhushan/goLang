// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // import "fmt"

// // func calculation(a, b int) int {
// // 	cal := a + b

// // 	return cal

// // }

// // func check(y []int, x func(a, b int) int) []int {
// // 	var results []int

// // 	for index, value := range y {
// // 		results = append(results, x(index, value))
// // 	}
// // 	return results
// // }

// func individual(x string) (string, string) {
// 	n := strings.ToUpper(x)

// 	names := strings.Split(n, " ")

// 	var initials []string

// 	for _, value := range names {

// 		initials = append(initials, value[:1])

// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}

// 	return "none", "none"

// }
// func main() {

// 	x, y := individual("hemanth bhushan")

// 	fmt.Println(x, y)
// }
