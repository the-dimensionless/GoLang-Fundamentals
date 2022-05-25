package main

import "fmt"

func main() {
	// method 1 (define type of key & value)
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#ffffff",
		"white": "#000000",
	}
	fmt.Println(colors)

	// method 2
	var c map[string]string
	fmt.Println(c)

	// method 3
	cols := make(map[string]string)
	cols["white"] = "#000000"
	fmt.Println(cols)

}
