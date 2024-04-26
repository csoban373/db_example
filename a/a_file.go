package a

import "fmt"

func PublicMethod() {
	fmt.Println("access outside of package a!")
}

type Person struct {
	Name   string
	credit int
}
