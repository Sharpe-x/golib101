package main

import (
	"excel/quick_start/example"
)

func main() {
	name := "Book1.xlsx"
	err := example.Create(name)
	if err != nil {
		panic(err)
	}

	err = example.Read(name)
	if err != nil {
		panic(err)
	}

}
