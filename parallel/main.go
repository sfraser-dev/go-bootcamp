package main

import (
	"fmt"
	"net/http"
)

func main() {
	var websites = []string{
		"https://google.com",
		"https://www.facebook.com",
		"https://go.dev",
		"https://stackoverflow.com",
		"https://www.amazon.co.uk",
	}

	// make creates a value out of the given type
	c := make(chan string)

	for _, website := range websites {
		//checkLink(website)
		go checkLink(website, c)
	}

	// pause and wait for channel data
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checkLink(website string, c chan string) {
	_, err := http.Get(website)
	if err != nil {
		msg := fmt.Sprintf("%s: %s", website, err)
		c <- msg
		return
	}
	// io.Copy(os.Stdout, resp.Body)
	msg := fmt.Sprintf("%s: yip, it's up", website)
	c <- msg
}
