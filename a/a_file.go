package a

import "fmt"

func PublicMethod() {
	fmt.Println("access outside of package a!")
}

func privateMethod() {
	fmt.Println("access inside package a!")
}

type Person struct {
	Name   string
	credit int
}
