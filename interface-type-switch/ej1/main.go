package main

import "fmt"

type I1 interface {
	M1()
}

type T1 struct{}

func (T1) M1() {}

type T2 struct{}

func (T2) M1() {}

func main() {
	var v I1 = T1{}
	switch aux := 1; v.(type) {
	case nil:
		fmt.Println("nil")
	case T1:
		fmt.Println("T1", aux)
	case T2:
		fmt.Println("T2", aux)
	}
}
