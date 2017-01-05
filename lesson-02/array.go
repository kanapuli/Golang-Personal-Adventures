package array

import (
	"fmt"
)

// Sometimes  we  need  to store a  collection of  items  with a   fixed length.
//Golang  provides  array for us
//Go's array  fixed size collection  of   single datatype  .

func ArrayDeclaration() {
	//An  array  can  be  declared using a var  keyword
	// The syntax  format is  var arrayName [arrayLength] datatype
	var myFirstArray [5]int

	myFirstArray[0] = 1
	myFirstArray[1] = 2
	//Fo other  unassigned positions ,  the default values  of the   datatype  will  come

	//Sometimes array  can  be  initialized in  the  declaration  part
	mySecondArray := [3]int{5, 6, 7}

	for _, item := range myFirstArray {
		fmt.Println(item)
	}

	for _, item = range mySecondArray {
		fmt.Println(item)
	}
}
