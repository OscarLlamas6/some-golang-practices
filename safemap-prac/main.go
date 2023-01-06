package main

import (
	"fmt"
	"safemap-prac/structs"
	"sync"
)

// Main func
func main() {
	//UsingSafeMap()

	sm := structs.New[int, int]()

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			sm.Insert(i, 6+i)

			value, ok := sm.Get(i)
			if !ok {
				fmt.Printf("El key con valor: %v no existe :o\n", value)
			}

			fmt.Printf("Se esperaba %v y se obtuvo %v\n", (6 + i), value)

		}(i)
	}
	wg.Wait()
}

// Just jusing safeMap function :P
func UsingSafeMap() {

	mySafeMap := structs.New[string, int]()
	mySafeMap.Insert("Money", 666)
	mySafeMap.Insert("Lines", 66611)

	value1, ok := mySafeMap.Get("Money")

	if ok {
		fmt.Println(value1)
	} else {
		fmt.Println("Value not found")
	}

	value2, ok := mySafeMap.Get("Money2")

	if ok {
		fmt.Println(value2)
	} else {
		fmt.Println("Value not found")
	}

	if updated := mySafeMap.Update("Money", 777); updated {
		fmt.Println("key value updated")
	} else {
		fmt.Println("key not found")
	}

	if updated := mySafeMap.Update("Money2", 777); updated {
		fmt.Println("key value updated")
	} else {
		fmt.Println("key not found")
	}

	if deleted := mySafeMap.Delete("Money"); deleted {
		fmt.Println("key value deleted succesfully")
	} else {
		fmt.Println("key not found")
	}

	if keyExists := mySafeMap.HasKey("Money"); keyExists {
		fmt.Println("Yeah! Key exists :D")
	} else {
		fmt.Println("F")
	}

	if keyExists := mySafeMap.HasKey("Lines"); keyExists {
		fmt.Println("Yeah! Key exists :D")
	} else {
		fmt.Println("F")
	}

}
