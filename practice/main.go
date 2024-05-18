package main

import (
	"fmt"
)

const aConst int = 64

func main() {

	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("the type is %T \n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "Hello"
	fmt.Println(anotherString)
	fmt.Printf("the type is %T \n", anotherString)

	anything := "Hi gg"
	fmt.Println(anything)
	fmt.Printf("the type is %T \n", anything)

	fmt.Println(aConst)
}
