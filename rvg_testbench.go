package main

import (
	"ISYE6644/project/rvg"
	"flag"
	"fmt"
	"strings"
)

const (
	Seed = 42
)

func main() {

	pDist := flag.String("dist", "uniform", "The distribution for which to generate RVs")
	pNumRVs := flag.Int("RVs", 1000, "The number of RVs to generate")
	var params rvg.Parameters
	flag.Var(&params, "params", "The parameters for the chosen distribution")

	flag.Parse()

	println(*pDist)
	println(*pNumRVs)
	fmt.Printf("%v\n", params)

	rv := make([]float64, *pNumRVs)

	var err error

	generator, exists := rvg.Generators[strings.ToLower(*pDist)]

	if !exists {
		println("The selected distribution is not available.")
		return
	}

	for i := range rv {
		rv[i], err = generator(params)
		if err != nil {
			println(err.Error())
			return
		}
	}

	rvg.WriteData(fmt.Sprintf("%s_%v.csv", *pDist, params), rv[:])

}
