package main

import "fmt"

type Interface interface {
	hello () string
}

type Test struct {}

func (t *Test) hello() string {
	return "hello"
}

func main() {
	fmt.Println("HELLO WORLD")
	blah := &Test{}
	fmt.Println(blah.hello())
}


