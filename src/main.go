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
	testBunc(blah)
}

func testFunc(b blah) {
	b.hello()
}

func testBunc(b blah) {
	b.hello()
}


