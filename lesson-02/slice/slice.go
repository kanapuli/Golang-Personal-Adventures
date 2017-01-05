package slice

import "fmt"

//array is used less in  golang .  Golang provides a new datatype built on top of 'array'type called slice
//slices are arrays without limit .This allows to store elements with dynamic length.

//An slice has  a  length and capacity .
//Slice Length is  number of  elements present in the slice
//Slice Capacity  is the number of elements  for which the space has  been allocated. So consider as max length
//Slice length can never exceed  the  capacity

//SliceDeclaration - Teaches you  how  to play with slices
func SliceDeclaration() {
	//Declaration  of Slice is  as simple as declaring an array without  giving the  length . Lets declare a nil slice
	var myFirstSlice []int

	//Disciple: Monk ! How to initialize a slice
	//Monk : 'Make' Slice  . initialize a slice using a make function
	//Monk :  myFirstSlice is declare  but if you try to assign values  , runtime  error will happen.
	//Monk :  First it should be initialized using make function
	myFirstSlice = make([]int, 3, 5)
	//So a Slice make  function should have a length and capacity . here length  = 3 and capacity = 5
	for _, element := range myFirstSlice {
		fmt.Println(element)
	}
	myFirstSlice = make([]int, 3)
	//Here length =capacity  = 3
	//Disciple :  Ok Monk Iam  clear about the make  functions. How will I assign values  to the slice
	//Monk : Lemme  show  you
	myFirstSlice[0] = 10
	myFirstSlice[1] = 4
	myFirstSlice[2] = 78
	for _, element := range myFirstSlice {
		fmt.Println(element)
	}

	//Monk: Yo Disciple ! There  is  another  way to create a slice . That is by using a slice literal
	//This is similar to an array literal
	sliceLiteral := []int{3, 4, 5} //  So here the length and capacity is 3
	for _, literal := range sliceLiteral {
		fmt.Println(literal)
	}
	//We can specify values to a particular index . Here the slice capacity become 20
	sliceLiteral = []int{0: 10, 20: 90}
	for _, literal := range sliceLiteral {
		fmt.Println(literal)
	}
	//Enlarging  Slices  using  Copy and  Append  function
	//Since  Slices  are dynamic  arrays  ,  we  can enlarge or  shrink the  slices
	//Create  two slices
	x := []int{10, 20, 30}
	fmt.Printf("Slice x Length is  %d and  capacity is  %d \n", len(x), cap(x))
	//Create  another slice y  where  the  length is 5  and  capacity is  10
	y := make([]int, 5, 10)
	//Call  the copy function  .  copy x  to y . syntax copy(destination, source)
	copy(y, x)
	fmt.Printf("Slice y Length is  %d and  capacity is  %d \n", len(y), cap(y))
	fmt.Println("After Copying , the  contents  of  y are : ", y)
	y[3] = 100
	y[4] = 500
	fmt.Println("After adding elements ,  the contents  of y are \n ,", y)

}
