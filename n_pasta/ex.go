package ex

import (
	"errors"
	"fmt"
)

func main() {
	a, b, err := nome("wesley", 11)
	if err != nil {
		panic("panico")
	}
	fmt.Println(a,b)
}

func nome(a string, b int) (string, int, error) {
	if b < 10 {
		return "", 0, errors.New("Deu ruim")
	}
	return a, b, nil
}