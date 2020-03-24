package main

import "fmt"

func main() {

	// Simple way of declating a variable
	var ironman string = "I am Ironman"
	fmt.Println(ironman)

	// Declaring variables using colon symbol (it's useful when the data-type is unclear to us)
	superman := "I am Superman"
	fmt.Println(superman)

	//knowing the data-type the has been assigned to variable uning "Printf"
	spiderman := "I am Spiderman"
	fmt.Printf("%v, %T", spiderman, spiderman) // "%v" will pass the value, & "%T" will show the data-type (like string/interger/etc..)

	//declaring multiple variable
	batman, joker := "I am Batman", "I am Joker"
	fmt.Println(batman, joker)

	// Declaring multiple variables unsing variable as a function
	var (
		wonderWoman = "I'm wonderWoman"
		health      = 100
		speed       = 750.86
	)
	fmt.Println(wonderWoman, health, speed)
}
