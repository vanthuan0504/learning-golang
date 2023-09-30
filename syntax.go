package main

import "fmt"

func main() {

	// var nameOne string = "one"
	// var nameTwo = "two"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "oneone"
	// nameThree = "three"
	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "stranger" // Can't use outside of a function
	// fmt.Println(nameFour)

	// ints
	// var ageOne int = 20
	// var ageTwo = 35
	// ageThree := 50
	// fmt.Println(ageOne, ageTwo, ageThree)

	// // bits & memory
	// var numOne int8 = 60
	// var numTwo int8 = -128
	// var numThree uint8 = 25
	// fmt.Println(numOne, numTwo, numThree)

	// var scoreOne float32 = 39.79
	// var scoreTwo float64 = 79472949147147198431
	// scoreThree := 33.33

	// fmt.Println(scoreOne, scoreTwo, scoreThree)

	// age := 35
	// name := "Shaun"
	// fmt.Print("Hello, ")
	// fmt.Print("world!")
	// fmt.Println("Your age is ", age)
	// fmt.Println("Your name is ", name)

	// // %_ = format specifier
	// fmt.Printf("Your age is %v and your name is %v\n", age, name)
	// fmt.Printf("Your age is %q and your name is %q\n", age, name)
	// fmt.Printf("Type of age is %T and Type of name is %T\n", age, name)
	// fmt.Printf("you scored %f points!\n", 225.55)
	// fmt.Printf("you scored %0.1f points!\n", 225.55)

	// // Sprintf (save formatted strings)
	// var str = fmt.Sprintf("Type of age is %v and Type of name is %v\n", age, name)
	// fmt.Println("the saved string is: ", str)

	// var ages [3]int = [3]int{30, 50, 100}
	var ages = [3]int{30, 50, 100}

	names := [4]string{"conan", "mario", "peach"}
	names[1] = "yasuo"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo, rangeThree)

}
