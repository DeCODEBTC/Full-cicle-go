package main

import (
	"fmt"
	"net/http"
)

func main() {

	res, err := http.Get("http://google.com.br")
	if err != nil {
		panic("Erro!")
	}

	fmt.Println(res.Status)
}
