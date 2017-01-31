package controlflow

import "fmt"

//DecisonStatements - Implements How If statements work in Go
func DecisonStatements() {
	//The decision statements in go are if , if else and  tf -else if - else

	//The if blocks gets executed if the  condition is true
	//Go syntax doesnt need a bracket around the  conditions
	if true {
		fmt.Println("This is going to be  executed")
	}

	if false {
		fmt.Println("Control will not come  inside this block")
	}
	//If statement using a negation operator
	if !false {
		fmt.Println("This also runs since the condition evalutaes to true")
	}

	//We can  do  initialization in an if statements
	if msg := "A value is initialized here"; true {
		fmt.Println(msg)
	}
	//This initialization is useful in case as follows ,
	//Printing errors
	/*if error := file.Chmod(0664); error != nil {
		log.Print(error)
	}*/

	//if else statement example
	nationality := "Indian"
	if nationality == "Antartican" {
		fmt.Println("This is not going to run")
	} else {
		fmt.Println("The nationality is Indian.")
	}
	//if -else if -else case
	if false {
		fmt.Println("No Control here")
	} else if true {
		fmt.Println("this is going to be executed")
	} else {
		fmt.Println("The control does not reach this block")
	}
}
