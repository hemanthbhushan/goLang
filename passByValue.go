package main

import "fmt"

func updateValue(x *string) {
	*x = "change"
}

func main() {

	name := "hemanth"

	fmt.Println("name location", &name)

	store := &name
	fmt.Println("name location in store", store)

	fmt.Println("name in pointer store", *store)

	fmt.Println("value in name before using pointer to update", name)

	updateValue(store)

	fmt.Println("value in name after using pointer to update", name)

}
