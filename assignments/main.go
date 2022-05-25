package main

import "fmt"

/* Create a slice of ints from 0 through 10
Iterate through the slice... for each element check if
element is even or odd */
func assignment1() {
	slice := []int{}
	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}

	for i, item := range slice {
		fmt.Printf("Index %d | item %d | type %s \n", i, item, test(item))
	}

}

func test(b int) string {
	if b%2 != 0 {
		return "Odd"
	} else {
		return "Even"
	}
}

func main() {
	assignment1()
}
