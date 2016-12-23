package main

import (
	"fmt"
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
}

//manipulateValue accepts a function as  a parameter
func manipulateValues(f func(input int) (int, int)) {
	c, d := f(5000000)
	fmt.Println(c, d)

	c, d = f(5798666)
	fmt.Println(c, d)
}
