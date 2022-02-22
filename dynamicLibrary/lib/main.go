package main

import "go-dynamicLibrary/def"

type EmptyDriver struct {
}

func (ed EmptyDriver) Name() string {
	return "EmptyDriver"
}

func NewDriver() def.Driver {
	return EmptyDriver{}
}
