package main

import "fmt"

func main() {
	sliceLen := 11
	mySlice := make([]int, sliceLen)
	for i := 0; i < sliceLen; i++ {
		mySlice[i] = i
	}
	fmt.Println(mySlice)
	for _, e := range mySlice {
		if e%2 == 0 {
			fmt.Printf("%02v: %s\n", e, "even")
		} else {
			fmt.Printf("%02v: %s\n", e, "odd")
		}
	}
}
