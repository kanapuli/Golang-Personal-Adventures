package hashtable

import "fmt"

//HashDeclaration : go's map function is similar to java's hashtable . It allows to store the data in key valuepairs
//HashDeclaration has the basic map type declarations
func HashDeclaration() {

	//The map declaration synta goes like this ,  map[key type]value type
	var myMap map[int]string
	//Now my map is a  nil  map. Attempting to write a value to a nil map will  genereate  runtime error.
	//It should be initalized using map function
	myMap = make(map[int]string)
	myMap[1] = "This"
	myMap[2] = "is"
	myMap[3] = "the map declaration"
	//Map won't allow duplicate keys
	myMap[3] = "Duplicate key"
	fmt.Println(myMap[3])
	language := map[string]string{
		"EN-IN": "English Britain",
		"EN-US": "English USA",
		"FR":    "French",
		//	"FR":    "French", Duplicate Key
	}
	fmt.Println(language)
	fmt.Println(language["FR"])

	/*Maps  provide  fastre  lookup  of  the  data  A lookup can be  performed by the  following way .*/
	if lng, ok := language["EN-IN"]; ok {
		fmt.Println(lng)
	}
	if lng, ok := language["ESP"]; ok {
		fmt.Println(lng)
	} else {
		fmt.Println("Dictionary key not found")
	}
}
