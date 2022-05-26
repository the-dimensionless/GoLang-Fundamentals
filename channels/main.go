package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link) // Blocking function call
	if err != nil {
		fmt.Printf("\n%v link might be down \n", link)
		return
	}
	fmt.Printf("%v link is up \n", link)
}
