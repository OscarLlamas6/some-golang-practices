package main

import "fmt"

type (
	Device interface {
		GetType() string
		GetRAMMemoryBrand() string
	}

	Computer struct {
		DeviceType string
		RamMemory
	}

	RamMemory struct {
		Value float64
		Brand string
	}
)

// Returns device type using Computer fields
func (m *Computer) GetType() string {
	return m.DeviceType
}

// Returns device RAM memory brand using struct composition
func (c *Computer) GetRAMMemoryBrand() string {
	return c.Brand
}

func NewComputer() *Computer {
	return &Computer{DeviceType: "Computer", RamMemory: RamMemory{
		Value: 16,
		Brand: "Kingston",
	}}

}

func main() {
	var device Device = NewComputer()

	fmt.Println("Device Type: ", device.GetType())
	fmt.Println("Device RAM Memory Brand: ", device.GetRAMMemoryBrand())

}
