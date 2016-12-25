package array

import (
	"fmt"
	"strconv"
)

// Sometimes  we  need  to store a  collection of  items  with a   fixed length.
//Golang  provides  array for us
//Go's array  fixed size collection  of   single datatype  .

//ArrayDeclaration  teaches about the golang's method of array declaration
func ArrayDeclaration() {
	//An  array  can  be  declared using a var  keyword
	// The syntax  format is  var arrayName [arrayLength] datatype
	var myFirstArray [5]int

	myFirstArray[0] = 1
	myFirstArray[1] = 2
	myFirstArray[2] = 3
	//Rest of the explicitly  unassigned array indexes will be  filled with datatype default values

	//Sometimes array  can  be  initialized in  the  declaration  part
	// 'new'  keyword is not required unlike in C# or Java
	mySecondArray := [3]int{5, 6, 7}

	//The array items  can be  dispalyed  by  running  a  for  loop
	for _, item := range myFirstArray {
		fmt.Println(item)
	}

	for _, item := range mySecondArray {
		fmt.Println(item)
	}
	//Golang  gives the liberty  to   specify the index and  assign  values
	myThirdArray := [5]string{0: "Zen", 1: "is", 2: "Acheivable"}
	//The array index 3 and  4  will   be   assignes  with  empty string values
	for index, item := range myThirdArray {
		if item == "" {
			//Here is an magic. index is an int type . So it cannot be  concatenated  with  string.  So strconv.Itoa is  used to convert int  to string .
			fmt.Println("The array index " + strconv.Itoa(index) + " has an  empty value.")
		} else {
			fmt.Println(item)
		}

	}
	//Array Literal with [...] specifies array of variable length
	myFourthArray := [...]int{23, 56, 89, 90, 89}
	fmt.Println(len(myFourthArray))
}
