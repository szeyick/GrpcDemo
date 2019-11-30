package models

import (
	"fmt"
)

type Foo interface {
	Bar(x int) int
}

type Impl struct {}

func (i *Impl) Bar(x int) int {
	fmt.Println("LO")
	return x
}

