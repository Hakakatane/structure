package structure

import (
	"fmt"
	"example.com/building"
)

// Здание
type Structure struct {
	Building  building.Building
	Area      float64
	Floors    int
	Address   string
	Condition string // эксплуатируется, горит, затоплено, разрушено, реконструкция
}

// Конструктор для Structure
func NewStructure(b building.Building, area float64, floors int, address string) Structure {
	return Structure{
		Building:  b,
		Area:      area,
		Floors:    floors,
		Address:   address,
		Condition: "эксплуатируется",
	}
}

// Методы для изменения состояния
func (s *Structure) Exploit() {
	s.Condition = "эксплуатируется"
	fmt.Println("Здание эксплуатируется.")
}

func (s *Structure) Fire() {
	s.Condition = "горит"
	fmt.Println("В здании пожар!")
}

func (s *Structure) Flood() {
	s.Condition = "затоплено"
	fmt.Println("Здание затоплено!")
}

func (s *Structure) Destroy() {
	s.Condition = "разрушено"
	fmt.Println("Здание разрушено!")
}

func (s *Structure) Renovate() {
	s.Condition = "реконструкция"
	fmt.Println("В здании ведется реконструкция.")
}

func (s *Structure) SetCondition(condition string) {
	s.Condition = condition
	fmt.Printf("Состояние здания изменено на %s\n", condition)
}

// Методы для получения характеристик
func (s Structure) GetArea() float64 {
	return s.Area
}

func (s Structure) GetFloors() int {
	return s.Floors
}

func (s Structure) GetAddress() string {
	return s.Address
}

func (s Structure) GetCondition() string {
	return s.Condition
}
