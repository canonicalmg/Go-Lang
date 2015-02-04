package main

import "fmt"

type Mystruct struct {
	mystring   string
}

func main() {

	nonpointervar := "Test me"
	var thepointer *string = &nonpointervar
	fmt.Println("Pointer stuff: ",*thepointer, thepointer)

	thisstruct := Mystruct{"Print me!"}
	thisstring := "I'm a string"
	fmt.Println(thisstruct)
	fmt.Println(thisstring)
	//Note that thisstring has no data type but rather is infered to be a string. Whereas the structure Mystruct is defined as a string
}
