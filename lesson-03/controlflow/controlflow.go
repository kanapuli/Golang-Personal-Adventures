package controlflow

//Switchcase - Function that executes switch case in golang
func Switchcase(input interface{}) (result interface{}) {
	switch input {
	case "Golang":
		return "Golang - A language developed by Google."
	case "C#":
		return "C# - A language developed by Microsoft."
	case 10:
		return "10 - A numerical number."
	default:
		return 200
	}
}
