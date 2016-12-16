/*package  is used to organize your codes
  The package 'main' code is always converted to an executable when running go build
  This is  equivalent to the Java/C# namespaces
*/
package main

//import is used to import the standard libraries and third party packages.
import "fmt"

//The program always start running from  the main function.
func main() {
	//Package  fmt is the Standar library which helps in formatted input/output .  Similar to C's printf and scanf
	//Println formats using the default format and  appends a new line character to the string by default
	fmt.Println("A journey of thousand  miles  starts with a single step")
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
	fmt.Printf("%s", "Become a GoLang Monk. \n")
}
