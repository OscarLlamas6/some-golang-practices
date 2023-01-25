package main

import "fmt"

type I1 interface {
	M1()
}

type T1 struct{}

func (T1) M1() {}

type I2 interface {
	I1
	M2()
}

type T2 struct{}

func (T2) M1() {}
func (T2) M2() {}

func main() {
	var v I1 = T2{}
	switch v.(type) {
	case I2:
		fmt.Println("I2")
	case T1:
		fmt.Println("T1")
	case T2:
		fmt.Println("T2")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("default")
	}
}
