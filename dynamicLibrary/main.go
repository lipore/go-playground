package main

import (
	"fmt"
	"go-dynamicLibrary/def"
	"plugin"
)

func main() {
	p, err := plugin.Open("driver.so")
	if err != nil {
		panic(err)
	}

	newDriverSymbol, err := p.Lookup("NewDriver")
	if err != nil {
		panic(err)
	}

	newDriverFunc := newDriverSymbol.(func() def.Driver)
	newDriver := newDriverFunc()
	fmt.Println(newDriver.Name())
}
