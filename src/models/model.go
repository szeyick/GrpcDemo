package models

type Foo interface {
	Bar(x int) int
}

type Impl struct {}

func (i *Impl) Bar(x int) int {
	return x
}

