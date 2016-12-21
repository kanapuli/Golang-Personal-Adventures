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

//Go Functions  can return multiple values and  the  return types  are specified inside the parantheses after the parameter lists
//SwapVariables : Now Swapping vaiable is as  simple as this .
func SwapVariables(x, y string) (string, string) {
	return y, x
}

//Example for Naked Return
//Naming Return Values  and concatenation of strinf
func NameReturnValues(x, y string) (result string) {
	result = x + " " + y
	//When the return value is named ,  then just the return statement is enough .  That automatically returns z . This is naked return
	return
}
