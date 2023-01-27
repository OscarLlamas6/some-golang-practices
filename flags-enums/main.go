package main

import "fmt"

const (
	Bold      = 1 << iota // 1 (001)
	Italic                // 2 (010)
	Underline             // 4 (100)
)

type Bitwise int

func (set Bitwise) Has(flag Bitwise) bool {
	return set&flag == flag
}

func (set Bitwise) Remove(flag Bitwise) Bitwise {
	return set &^ flag
}

func (set Bitwise) Add(flag Bitwise) Bitwise {
	return set | flag
}

func main() {
	var formatTypes Bitwise = Underline | Italic
	fmt.Println(formatTypes) // 6: Underline | Italic
	formatTypes = formatTypes.Add(Bold).Remove(Underline)
	fmt.Println(formatTypes) // 3: Bold | Italic
}
