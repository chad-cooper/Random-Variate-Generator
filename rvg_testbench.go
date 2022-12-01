package main

import (
	"ISYE6644/project/rvg"
	"flag"
	"fmt"
	"strings"
	"sync"
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

	// Lowercase the distribution name
	*pDist = strings.ToLower(*pDist)

	fmt.Printf("Generating %d RVs for %s(%v) distribution.\n", *pNumRVs, *pDist, params)

	rv := make([]float64, *pNumRVs)

	generator, exists := rvg.Generators[*pDist]

	if !exists {
		println("The selected distribution is not available.")
		return
	}

	// Validate params
	if paramsInvalid := rvg.ParamsValidate(params, *pDist); paramsInvalid != nil {
		println(paramsInvalid.Error())
		return
	}

	// Set up for multi-goroutine generation
	var wg sync.WaitGroup

	wg.Add(len(rv))

	for i := range rv {
		go func(i int) {
			defer wg.Done()
			rv[i], _ = generator(params)
		}(i)
	}

	wg.Wait()

	rvg.WriteData(fmt.Sprintf("%s_%v.csv", *pDist, params), rv[:])

}
