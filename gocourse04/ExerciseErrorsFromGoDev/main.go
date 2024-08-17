package main

import (
	"fmt"
	"time"
)

type ErrNegativeSqrt struct {
	When time.Time
	What float64
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt{time.Now(), x}
	}
	return 0, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("at %v, cannot Sqrt negative number: %v",
		e.When, e.What)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
