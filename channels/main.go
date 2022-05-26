package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//fmt.Println(<-c) // blocking wait on channel for message
	for l := range c { //infinte loop
		go func(l string) {
			time.Sleep(5 * time.Second)
			go checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // Blocking function call
	if err != nil {
		fmt.Printf("\n %v link might be down \n", link)
		c <- " Might be down \n"
		return
	}
	fmt.Printf("\n %v link is up \n", link)
	c <- " Yep! Link is up \n"
}
