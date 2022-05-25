package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// m1
	alex := person{"Alex", "Hales", contactInfo{"alex@gmail.com", 12345}}
	// m2
	sumit := person{
		firstName: "Sumit",
		lastName:  "Singh",
		contact: contactInfo{
			email:   "sumit.03july@gmail.com",
			zipCode: 112345,
		}, // contact key is optional, only the type will work too!
	}
	fmt.Println(alex, sumit)

	// m3
	var tony person // default is ZERO VALUE
	tony.firstName = "Tony"
	tony.lastName = "Stark"
	tony.contact.email = "tony@stark.com"
	tony.contact.zipCode = 10880
	fmt.Printf("%+v", tony)
}
