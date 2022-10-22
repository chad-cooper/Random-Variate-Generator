package main

import (
	"ISYE6644/project/rvg"
	"errors"
	"fmt"
	"strconv"
	"strings"

	// "os"
	"flag"
)

const (
	Seed  = 42
	rvCap = 1000
)

type parameters []float64

func main() {

	pDist := flag.String("dist", "uniform", "The distribution for which to generate RVs")

	pNumRVs := flag.Int("RVs", 1000, "The number of RVs to generate")

	var params parameters

	flag.Var(&params, "params", "The parameters for the chosen distribution")

	flag.Parse()

	println(*pDist)
	println(*pNumRVs)
	fmt.Printf("%v\n", params)

	rv := make([]float64, *pNumRVs)

	// :TODO: error check for valid generator

	for i := range rv {
		rv[i] = rvg.Generators[strings.ToLower(*pDist)](params)
	}

	rvg.WriteData(fmt.Sprintf("%s_%v.csv", *pDist, params), rv[:])

}

func (p *parameters) String() string {
	return fmt.Sprintf("%v", *p)
}

func (p *parameters) Set(paramList string) error {
	if len(*p) > 0 {
		return errors.New("Parameter flag already set")
	}

	for _, param := range strings.Split(paramList, " ") {
		val, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return err
		}

		*p = append(*p, val)
	}
	return nil
}
