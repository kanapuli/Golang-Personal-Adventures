package main

import (
	"fmt"

	cf "github.com/Golang-The-Zen-Way-/lesson-03/controlflow"
)

func main() {
	//call the switch case methods
	fmt.Println("An example for Switch case control flow. No need to specify break statement ")
	result := cf.Switchcase("C#")
	fmt.Println(result)
	result = cf.Switchcase(500) //expects default statement here
	fmt.Println(result)
	result = cf.Switchcase(10)
	fmt.Println(result)
}
