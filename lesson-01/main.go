/*package  is used to organize your codes
  The package 'main' code is always converted to an executable code when running go build
  This is  equivalent to the Java/C# namespaces
*/
package main

//import is used to import the standard libraries and third party packages.
import (
	"fmt"
	basics "github.com/Golang-The-Zen-Way-/lesson-01/catchthebasics"
)

//The program always start running from  the main function.
func main() {
	//Package  fmt is the Standar library which helps in formatted input/output .  Similar to C's printf and scanf
	//Println formats using the default format and  appends a new line character to the string by default
	fmt.Println("A journey of thousand  miles  starts with a single step")

	//Call Another function from the function package from the file catchthebasics.go
	additonResult := basics.Add(100, 8)
	fmt.Println("The Sacred Number in zen is ", additonResult)

	//Call the SwapVariables  function from  the basics package
	//This is an example to show that GO Functions can return multiple variables
	a, b := basics.SwapVariables("Yang", "Yin")
	//So the SwapVariables should return results such that a = Yang and b = Yin
	fmt.Println("GoLang swaps ", a, " and ", b, "with ease . Follow Golang ! :)")
	//Call the NameReturnValues function
	truth := basics.NameReturnValues("Show me the truth", "Master")
	fmt.Println(truth)
	//Call the Variadic function from basics package
	zenPrinciples1 := basics.VariadicFunction("Right Speech", "Right Behaviour", "Right Action", "Right Thinking")
	zenPrinciples2 := basics.VariadicFunction("Right Attitude", "Right Contemplation")
	fmt.Println("Zen Principles Wont Vary. But The Parameter List can vary for a GOlang Variadic Function")
	fmt.Println(zenPrinciples1)
	fmt.Println(zenPrinciples2)

}

/*The init function runs  even before the main function.  The init function can also be in other packages. So when importing  a package which contains
an init function , then that function is run  first.
It is possible to have multiple init functions . They will be executed in the order of how they show up in code  */
func init() {
	//Printf uses the custom verbs  to format the string
	/*  %v -  default format of the string
	    %T -   Type of the string
	    %t -  Boolean
	    %b -  base 2
	    %d -  base 10
	    %o -  base 8
	    %U -  Unicode format
	    %e or  %E - Scientefic notation
	    %s -  String */
	fmt.Printf("%s", "Here are the Lessons that will make you a Golang Monk. \n")
}
