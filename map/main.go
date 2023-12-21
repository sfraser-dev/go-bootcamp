package main

import "fmt"

func main() {
	// declare a map literally
	colors_lit := map[string]string{
		"red":   "ff0000",
		"green": "4bb888",
	}
	fmt.Println(colors_lit)

	// declare a map using var
	var colors_var = map [string]int {
		"a": 1,
		"b": 2,
		"y": 25,
		"z": 26,
	}
	fmt.Println(colors_var)

	// declare a map by inferring
	colors_infer := map[string]bool{
		"1": true,
		"0": false,
	}
	fmt.Println(colors_infer)

	// declare a map using inbuild make function
	colors_make := make(map[string]float32)
	colors_make["price"] =  10.99
	colors_make["discountPercentage"] =  15
	fmt.Println(colors_make)
}
