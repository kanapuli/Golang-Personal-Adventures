package slice

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
	myFirstSlice := make([]int, 3, 5)
	//So a Slice make  function should have a length and capacity . here length  = 3 and capacity = 5
	myFirstSlice := make([]int, 3)
	//Here length =capacity  = 3
	//Disciple :  Ok Monk Iam  clear about the make  functions. How will I assign values  to the slice
	//Monk : Lemme  show  you
	myFirstSlice[0] = 10
	myFirstSlice[1] = 4
	myFirstSlice[2] = 78
}
