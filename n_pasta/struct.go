package structs

import "fmt"

type Car struct {
	CarName string
	CarYear int
}

func (c Car) drive() {
	fmt.Println(c.CarName, "Andou!")
}

func main() {
	car1 := Car{
		CarName: "Ferrari",
		CarYear: 2020,
	}

	car1.drive()
}