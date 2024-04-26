package a

import "fmt"

func PrintExtra() {
	fmt.Println("hi there")
}

func privateMethod() {
	fmt.Println("you can't access me outside of package a!")
}
