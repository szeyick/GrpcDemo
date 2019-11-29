package main

import "fmt"

type blah interface {
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
	testFunc(blah)
}

func testFunc(b blah) {
	b.hello()
}


