package rvg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Define Parameters type and implement Flags interface
type Parameters []float64

func (p *Parameters) String() string {
	return fmt.Sprintf("%v", *p)
}

func (p *Parameters) Set(paramList string) error {
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

// Parameter error types
type InvalidParametersError struct {
	dist           Distribution
	paramsSupplied Parameters
}

type NumberOfParametersError struct {
	dist           Distribution
	paramsSupplied Parameters
}

func (e *InvalidParametersError) Error() string {

	// s := "Invalid parameters supplied for the selected distribution.\nExpected: "
	s := fmt.Sprintf("Invalid parameters supplied for the %s distribution.\nExpected: ", e.dist.name)

	for _, lim := range e.dist.paramBounds {
		s += fmt.Sprintf("%s, ", lim)
	}

	s = s + "\nGot: "

	for i, sym := range e.dist.paramSymbols {
		s += fmt.Sprintf("%s = %.2f, ", sym, e.paramsSupplied[i])
	}

	return s
}

func (e *NumberOfParametersError) Error() string {

	numSupplied := len(e.paramsSupplied)

	var s string

	if numSupplied < e.dist.numParams {
		s = "Too few"
	} else {
		s = "Too many"
	}

	s = fmt.Sprintf("%s arguments were supplied for the %s(%s) distribution.", s, e.dist.name, strings.Join(e.dist.paramSymbols, ","))

	return fmt.Sprintf("%s\nExpected: %d\nGot: %d\n", s, e.dist.numParams, len(e.paramsSupplied))
}

type RVGParamsValidate func(Parameters) error

// Distribution details
type Distribution struct {
	name         string
	numParams    int
	paramSymbols []string
	paramBounds  []string
	validBounds  func(Parameters) bool
}

var DistDetails = map[string]Distribution{
	"bernoulli": {
		name:         "Bernoulli",
		numParams:    1,
		paramSymbols: []string{"p"},
		paramBounds:  []string{"0 ≤ p ≤ 1"},
		validBounds: func(params Parameters) bool {
			return 0 <= params[0] && params[0] <= 1
		},
	},
	"binomial": {
		name:         "Binomial",
		numParams:    2,
		paramSymbols: []string{"n", "p"},
		paramBounds:  []string{"n ∈ {0, 1, 2,...}", "0 ≤ p ≤ 1"},
		validBounds: func(params Parameters) bool {
			return (params[0] == float64(int(params[0]))) && (params[0] >= 0) && (0 <= params[1] && params[1] <= 1)
		},
	},
	"exponential": {
		name:         "Exponential",
		numParams:    1,
		paramSymbols: []string{"λ"},
		paramBounds:  []string{"λ > 0"},
		validBounds: func(params Parameters) bool {
			return params[0] > 0
		},
	},
	"gamma": {
		name:         "Gamma",
		numParams:    2,
		paramSymbols: []string{"k", "θ"},
		paramBounds:  []string{"k > 0", "θ > 0"},
		validBounds: func(params Parameters) bool {
			return (params[0] > 0) && (params[1] > 0)
		},
	},
	"geometric": {
		name:         "Geometric",
		numParams:    1,
		paramSymbols: []string{"p"},
		paramBounds:  []string{"0 < p ≤ 1"},
		validBounds: func(params Parameters) bool {
			return 0 < params[0] && params[0] <= 1
		},
	},
	"normal": {
		name:         "Normal",
		numParams:    2,
		paramSymbols: []string{"μ", "σ"},
		paramBounds:  []string{"μ ∈ ℝ", "σ > 0"},
		validBounds: func(params Parameters) bool {
			return params[1] > 0
		},
	},
	"poisson": {
		name:         "Poisson",
		numParams:    1,
		paramSymbols: []string{"λ"},
		paramBounds:  []string{"λ > 0"},
		validBounds: func(params Parameters) bool {
			return params[0] > 0
		},
	},
	"weibull": {
		name:         "Weibull",
		numParams:    2,
		paramSymbols: []string{"λ", "k"},
		paramBounds:  []string{"λ > 0", "k > 0"},
		validBounds: func(params Parameters) bool {
			return (params[0] > 0) && (params[1] > 0)
		},
	},
}

func paramsValidate(params Parameters, distName string) error {

	dist := DistDetails[distName]

	if len(params) != dist.numParams {
		return &NumberOfParametersError{
			paramsSupplied: params,
			dist:           dist,
		}
	}

	if !dist.validBounds(params) {
		return &InvalidParametersError{
			paramsSupplied: params,
			dist:           dist,
		}
	}

	return nil
}
