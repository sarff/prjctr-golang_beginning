package main

import (
	"github.com/brianvoe/gofakeit/v6"
)

// Create structs with random injected data
type Fake struct {
	Str     string
	Int     int
	Pointer *int
	Name    string `fake:"{firstname}"` // Any available function all lowercase
	Gender  string `fake:"{gender}"`
}

type Bar struct {
	Name   string
	Number int
	Float  float32
}

func GetFaker() Fake {
	var f Fake
	gofakeit.Struct(&f)

	return f
}
