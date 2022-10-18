package main

import (
	"ISYE6644/project/rvg"
)

const (
	Seed  = 42
	rvCap = 1000
)

func main() {

	rv := make([]float64, 1000, 1000)

	for i := 0; i < cap(rv); i++ {
		rv[i] = rvg.ExponentialGenerator(1.0)
	}

	rvg.WriteData("exponential_lambda1.csv", rv)

}
