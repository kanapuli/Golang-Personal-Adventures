package main

import (
	"fmt"

	hashTable "github.com/Golang-The-Zen-Way-/lesson-02/hashtable"

	array "github.com/GoLang-The-Zen-Way-/lesson-02/array"
	slice "github.com/Golang-The-Zen-Way-/lesson-02/slice"
)

func main() {
	a, b := 100, 200
	//Example  for Anonymous Function
	//In Golang a   Function  can  be  declared inside a  function
	//Here is an anonymous function which get one input value  and returns two values
	result := func(input int) (int, int) {
		x := input / (1000 + a)
		y := input * 99 * ((a + b) / b)
		//return the values
		return x, y
	}
	//In Golang ,  you can pass  a  function like a  value
	manipulateValues(result)

	/*  Lets  start  our  Array Lesson */
	//So  an array package is  imported .
	array.ArrayDeclaration()

	/*Lets learn about slice  */
	slice.SliceDeclaration()
	/*map declaration */
	hashTable.HashDeclaration()

}

//manipulateValue accepts a function as  a parameter
func manipulateValues(f func(input int) (int, int)) {
	c, d := f(5000000)
	fmt.Println(c, d)

	c, d = f(5798666)
	fmt.Println(c, d)
}
