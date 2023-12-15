package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode uint32
}

type person struct {
	firstName string
	lastName  string
	// shorthand for contactInfo contactInfo
	contactInfo
}

func (p person) printInfo() {
	fmt.Printf("%+v\n", p)
}

func (ptr *person) updateName(n string) {
	(*ptr).firstName = n
}

func main() {
	// attributes set via order or arguments
	greg := person{
		"Greg",
		"Norman",
		contactInfo{"greg@example.com", 12345},
	}

	fred := person{
		firstName: "Fred",
		lastName:  "Couples",
		contactInfo: contactInfo{
			"fred@example.com",
			54321},
	}
	
	// zero value initialisation of nick
	var nick person 
	fmt.Println(greg)
	fmt.Println(fred)
	fmt.Printf("%+v\n", nick)

	// fill out nick's info
	nick.firstName = "Nick"
	nick.lastName = "Faldo"
	nick.contactInfo.email = "nick@example.com"
	nick.contactInfo.zipCode = 13579
	nick.printInfo()

	nickPtr := & nick
	nickPtr.updateName("Nicholas")
	nick.printInfo()

}
