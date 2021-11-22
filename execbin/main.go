package main

import (
	"fmt"

	"github.com/goccy/go-execbin"
)

func main() {
	file, err := execbin.Open("./execbin")
	if err != nil {
		panic(err)
	}

	types, err := file.DefinedInterfaceTypes()
	if err != nil {
		panic(err)
	}
	fmt.Println(types[0])
}
