package catchthebasics

/* Writing Functions
Functions in Golang has the following structure
func name (list of paramas) optional_return_type
{
    function body
} */

/* Any function which starts with an Upper case is an exported function. It can be called
in other packages also. So basicallly , it is equivalent to a public method.
The exported function should always have comment starting with the method name */

//Add :  Adds two integers and returns an int
func Add(x, y int) int {
	return x + y
}
