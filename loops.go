// package main

// import "fmt"

// func main() {

// 	arr := []string{"what", "you", "want", "MF"}

// 	for i := 0; i < len(arr); i++ {
// 		fmt.Println("the strings in the arr are", arr[i])
// 	}

// 	for index, value := range arr {

// 		fmt.Printf("the index of arr %v and the value in array is %v \n", index, value)

// 	}

// 	//if we dont need index replace index with _

// 	for _, value := range arr {

// 		fmt.Printf("the value in array is %v \n", value)
// 		value = "happy"

// 	}

// 	fmt.Println("the arr value doesnt change even if we update the value cuz its doesnt update in the value in ths slice it just creates the copy", arr)

// }
