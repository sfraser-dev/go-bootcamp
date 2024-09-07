package main

import "fmt"

func main() {
	// declare a map literally
	colorsLiteral := map[string]string{
		"red":   "#ff0000",
		"green": "#4bb888",
		"white": "#ffffff",
	}
	fmt.Println(colorsLiteral)

	// declare a map using var (variable declaration and initialization)
	var colorsVar = map[string]int{
		"a": 1,
		"b": 2,
		"y": 25,
		"z": 26,
	}
	fmt.Println(colorsVar)

	// declare a map by inferring (short variable declaration)
	colorsInfer := map[string]bool{
		"1": true,
		"0": false,
	}
	fmt.Println(colorsInfer)

	// declare a map using inbuilt make function
	colorsMake := make(map[string]float32)
	colorsMake["price"] = 10.99
	colorsMake["discountPercentage"] = 15
	fmt.Println(colorsMake)

	// make and delete for maps
	drinks := make(map[int]string)
	drinks[10] = "Pepsi"
	fmt.Println(drinks)
	delete(drinks, 10)
	fmt.Println(drinks)

	// passing map as argument to function
	mapPrint(colorsLiteral)
}

// map is a reference type, so true source of data changed
func mapPrint(m map[string]string) {
	for k, v := range m {
		fmt.Print("key=", k, ", value=", v, "\n")
	}
}
