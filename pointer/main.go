package main

import "fmt"

func updateName(x string) string {
	x = "yasuo"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99

}

func updateNameValue(n *string) {
	*n = "yasuo"
}

func main() {

	// group A types -> strings, ints, bools, floats, arrays, structs
	// non-pointer wrapper values
	name := "tifa"
	name = updateName(name)

	fmt.Println(name)

	// groups B types => slices, maps, functions
	// pointers wrapper values
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)

	fmt.Println("--------------memory address -----------")

	name = "tifa"
	// & gets the memory address of the value (pointer)
	fmt.Println("Memory address of name is: ", &name)

	// * gets the value at the specified memory address
	m := &name
	fmt.Println("Memory address: ", m)
	fmt.Println("Value at memory address: ", *m)

	updateName(*m) // using pointer as arg, can pass &name directly instead of "m" as well
	fmt.Println(name)

}
