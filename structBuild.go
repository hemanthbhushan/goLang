package main

import "fmt"

type b struct {
	name   string
	rollNo uint
	marks  map[string]uint
}

func newBill() b {

	bill := b{
		name:   "hemanth",
		rollNo: 23,
		marks:  map[string]uint{"telugu": 23, "hindi": 34},
	}

	return bill
}

func (bill b) structFunction() string {

	for k, v := range bill.marks {

		fmt.Println(k, v)

	}
	return "completed"

}
