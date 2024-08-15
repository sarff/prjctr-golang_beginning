package main

import (
	"github.com/brianvoe/gofakeit/v7"
)

type Fake struct {
	Name   string `fake:"{firstname}"` // Any available function all lowercase
	Gender string `fake:"{gender}"`
}

func NewFake() Fake {
	var f Fake
	gofakeit.Struct(&f)

	return f
}
