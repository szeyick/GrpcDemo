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
	Sum(5, 5)
}

func testFunc(b blah) {
	b.hello()
}

func testBunc(b blah) {
	b.hello()
}

func Sum(x int, y int) int {
	return x + y
}


