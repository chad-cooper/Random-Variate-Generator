package rvg

import (
	"fmt"
	"math"
	"testing"
)

// Test
func TestBernoulliGenerator(t *testing.T) {

	var tests = []struct {
		params  Parameters
		wantErr error
	}{
		{Parameters{0.5}, nil},
		{Parameters{0}, nil},
		{Parameters{1}, nil},
		{Parameters{-1.0}, &InvalidParametersError{}},
		{Parameters{1.5}, &InvalidParametersError{}},
		{Parameters{}, &NumberOfParametersError{}},
		{Parameters{1, 2}, &NumberOfParametersError{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Bernoulli{%v}", &tt.params)
		t.Run(testname, func(t *testing.T) {
			ans, err := BernoulliGenerator(tt.params)

			errType := fmt.Sprintf("%T", err)
			wantErrType := fmt.Sprintf("%T", tt.wantErr)

			if errType != wantErrType {
				t.Errorf("Got %T, wanted %T", err, tt.wantErr)
			}

			if err != nil {
				println(err.Error())
				return
			}

			if !(ans == 0 || ans == 1) {
				t.Errorf("Got %f, wanted either 0 or 1", ans)
			}

		})
	}

}

func TestBinomialGenerator(t *testing.T) {

	var tests = []struct {
		params  Parameters
		wantErr error
	}{
		// Standard cases
		{Parameters{3, 0.5}, nil},
		{Parameters{0, 0.5}, nil},
		{Parameters{1000, 0.5}, nil},
		{Parameters{3, 0}, nil},
		{Parameters{3, 1}, nil},
		// Invalid parameter cases
		{Parameters{-1, 0.5}, &InvalidParametersError{}},
		{Parameters{1.6, 0.5}, &InvalidParametersError{}},
		{Parameters{3, -1}, &InvalidParametersError{}},
		{Parameters{3, 2}, &InvalidParametersError{}},
		// Incorrect number of parameters cases
		{Parameters{}, &NumberOfParametersError{}},
		{Parameters{3}, &NumberOfParametersError{}},
		{Parameters{3, 0.5, 1}, &NumberOfParametersError{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Binomial{%v}", &tt.params)
		t.Run(testname, func(t *testing.T) {
			ans, err := BinomialGenerator(tt.params)

			errType := fmt.Sprintf("%T", err)
			wantErrType := fmt.Sprintf("%T", tt.wantErr)

			if errType != wantErrType {
				t.Errorf("Got %T, wanted %T", err, tt.wantErr)
			}

			if err != nil {
				println(err.Error())
				return
			}

			// Check if ans is integer
			if ans != float64(int(ans)) {
				t.Errorf("RV should be integer. Recieved RV = %.2f", ans)
			}

			// Check that ans is in proper range
			if !(0 <= ans && ans <= tt.params[0]) {
				t.Errorf("RV should be in range [0, %.2f], but it was %.2f.", tt.params[0], ans)
			}

		})
	}

}

func TestExponentialGenerator(t *testing.T) {

	var tests = []struct {
		params  Parameters
		wantErr error
	}{
		// Standard cases
		{Parameters{1}, nil},
		{Parameters{2.5}, nil},
		{Parameters{0.0001}, nil},
		{Parameters{1000}, nil},

		// Invalid parameter cases
		{Parameters{0}, &InvalidParametersError{}},
		{Parameters{-1.5}, &InvalidParametersError{}},
		// Incorrect number of parameters cases
		{Parameters{}, &NumberOfParametersError{}},
		{Parameters{1, 2}, &NumberOfParametersError{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Exponential{%v}", &tt.params)
		t.Run(testname, func(t *testing.T) {
			ans, err := ExponentialGenerator(tt.params)

			errType := fmt.Sprintf("%T", err)
			wantErrType := fmt.Sprintf("%T", tt.wantErr)

			if errType != wantErrType {
				t.Errorf("Got %T, wanted %T", err, tt.wantErr)
			}

			if err != nil {
				println(err.Error())
				return
			}

			if ans < 0 {
				t.Errorf("RV should be greater than or equal to 0, but it was %.2f.", ans)
			}

		})
	}

}

func TestGammaGenerator(t *testing.T) {

	var tests = []struct {
		params  Parameters
		wantErr error
	}{
		// Standard cases
		{Parameters{1, 1}, nil},
		{Parameters{0.5, 1}, nil},
		{Parameters{2, 1}, nil},
		{Parameters{1, 0.5}, nil},
		{Parameters{0.5, 0.5}, nil},
		{Parameters{2, 0.5}, nil},
		{Parameters{1, 2}, nil},
		{Parameters{0.5, 2}, nil},
		{Parameters{2, 2}, nil},
		{Parameters{1, 0.0001}, nil},
		{Parameters{0.01, 1}, nil},
		{Parameters{10, 10}, nil},

		// Invalid parameter cases
		{Parameters{0, 1}, &InvalidParametersError{}},
		{Parameters{1, 0}, &InvalidParametersError{}},
		{Parameters{-1, 1}, &InvalidParametersError{}},
		{Parameters{1, -1}, &InvalidParametersError{}},
		// Incorrect number of parameters cases
		{Parameters{}, &NumberOfParametersError{}},
		{Parameters{1}, &NumberOfParametersError{}},
		{Parameters{1, 2, 3}, &NumberOfParametersError{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Gamma{%v}", &tt.params)
		t.Run(testname, func(t *testing.T) {
			ans, err := GammaGenerator(tt.params)

			errType := fmt.Sprintf("%T", err)
			wantErrType := fmt.Sprintf("%T", tt.wantErr)

			if errType != wantErrType {
				t.Errorf("Got %T, wanted %T", err, tt.wantErr)
			}

			if err != nil {
				println(err.Error())
				return
			}

			if ans <= 0 {
				t.Errorf("RV should be greater than 0, but it was %.2f", ans)
			}

		})

	}

}

func TestGeometricGenerator(t *testing.T) {

	var tests = []struct {
		params  Parameters
		wantErr error
	}{
		{Parameters{0.5}, nil},
		{Parameters{1}, nil},
		{Parameters{0}, &InvalidParametersError{}},
		{Parameters{-1.0}, &InvalidParametersError{}},
		{Parameters{1.5}, &InvalidParametersError{}},
		{Parameters{}, &NumberOfParametersError{}},
		{Parameters{1, 2}, &NumberOfParametersError{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Geometric{%v}", &tt.params)
		t.Run(testname, func(t *testing.T) {
			ans, err := GeometricGenerator(tt.params)

			errType := fmt.Sprintf("%T", err)
			wantErrType := fmt.Sprintf("%T", tt.wantErr)

			if errType != wantErrType {
				t.Errorf("Got %T, wanted %T", err, tt.wantErr)
			}

			if err != nil {
				println(err.Error())
				return
			}

			// Check if ans is integer
			if ans != float64(int(ans)) {
				t.Errorf("RV should be integer. Recieved RV = %.2f", ans)
			}

			// Check that ans is in proper range
			if !(0 <= ans && ans <= math.Ceil(tt.params[0])) {
				t.Errorf("RV should be in range [0, %.2f], but it was %.2f.", math.Ceil(tt.params[0]), ans)
			}

		})
	}

}

func TestNormalGenerator(t *testing.T) {

	var tests = []struct {
		params  Parameters
		wantErr error
	}{
		{Parameters{0, 1}, nil},
		{Parameters{0, 0.00001}, nil},
		{Parameters{0, 0}, &InvalidParametersError{}},
		{Parameters{0, -1.0}, &InvalidParametersError{}},
		{Parameters{}, &NumberOfParametersError{}},
		{Parameters{1}, &NumberOfParametersError{}},
		{Parameters{1, 2, 3}, &NumberOfParametersError{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Normal{%v}", &tt.params)
		t.Run(testname, func(t *testing.T) {
			ans, err := NormalGenerator(tt.params)

			errType := fmt.Sprintf("%T", err)
			wantErrType := fmt.Sprintf("%T", tt.wantErr)

			if errType != wantErrType {
				t.Errorf("Got %T, wanted %T", err, tt.wantErr)
			}

			if err != nil {
				println(err.Error())
				return
			}

			if !(-4*tt.params[1] < ans && ans < 4*tt.params[1]) {
				t.Errorf("RV should be in range (%.2f, %.2f) but got RV = %.2f", -4*tt.params[1], 4*tt.params[1], ans)
			}

		})
	}

}

func TestPoissonGenerator(t *testing.T) {

	var tests = []struct {
		params  Parameters
		wantErr error
	}{
		// Standard cases
		{Parameters{1}, nil},
		{Parameters{2.5}, nil},
		{Parameters{0.01}, nil},
		{Parameters{50}, nil},

		// Invalid parameter cases
		{Parameters{0}, &InvalidParametersError{}},
		{Parameters{-1.5}, &InvalidParametersError{}},
		// Incorrect number of parameters cases
		{Parameters{}, &NumberOfParametersError{}},
		{Parameters{1, 2}, &NumberOfParametersError{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Poisson{%v}", &tt.params)
		t.Run(testname, func(t *testing.T) {
			ans, err := PoissonGenerator(tt.params)

			errType := fmt.Sprintf("%T", err)
			wantErrType := fmt.Sprintf("%T", tt.wantErr)

			if errType != wantErrType {
				t.Errorf("Got %T, wanted %T", err, tt.wantErr)
			}

			if err != nil {
				println(err.Error())
				return
			}

			if ans < 0 {
				t.Errorf("RV should be greater than or equal to 0, but got RV = %.2f", ans)
			}

		})
	}

}

func TestTriangularGenerator(t *testing.T) {

	var tests = []struct {
		params  Parameters
		wantErr error
	}{
		// Standard cases
		{Parameters{0, 2, 1}, nil},
		{Parameters{-1, 1, 0}, nil},
		{Parameters{1, 5, 3}, nil},
		{Parameters{-10, -1, -5}, nil},

		// Invalid parameter cases
		{Parameters{0, -2, 1}, &InvalidParametersError{}},
		{Parameters{0, 2, -1}, &InvalidParametersError{}},
		// Incorrect number of parameters cases
		{Parameters{}, &NumberOfParametersError{}},
		{Parameters{0}, &NumberOfParametersError{}},
		{Parameters{0, 2}, &NumberOfParametersError{}},
		{Parameters{0, 2, 1, 0.5}, &NumberOfParametersError{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Triangular{%v}", &tt.params)
		t.Run(testname, func(t *testing.T) {
			ans, err := TriangularGenerator(tt.params)

			errType := fmt.Sprintf("%T", err)
			wantErrType := fmt.Sprintf("%T", tt.wantErr)

			if errType != wantErrType {
				t.Errorf("Got %T, wanted %T", err, tt.wantErr)
			}

			if err != nil {
				println(err.Error())
				return
			}

			if ans < tt.params[0] || ans > tt.params[1] {
				t.Errorf("RV should be between %.2f and %.2f, but got RV = %.2f", tt.params[0], tt.params[1], ans)
			}

		})
	}

}

func TestWeibullGenerator(t *testing.T) {

	var tests = []struct {
		params  Parameters
		wantErr error
	}{
		// Standard cases
		{Parameters{1, 1}, nil},
		{Parameters{0.5, 1}, nil},
		{Parameters{2, 1}, nil},
		{Parameters{1, 0.5}, nil},
		{Parameters{0.5, 0.5}, nil},
		{Parameters{2, 0.5}, nil},
		{Parameters{1, 2}, nil},
		{Parameters{0.5, 2}, nil},
		{Parameters{2, 2}, nil},
		{Parameters{1, 0.0001}, nil},
		{Parameters{0.0001, 1}, nil},
		{Parameters{10, 10}, nil},

		// Invalid parameter cases
		{Parameters{0, 1}, &InvalidParametersError{}},
		{Parameters{1, 0}, &InvalidParametersError{}},
		{Parameters{-1, 1}, &InvalidParametersError{}},
		{Parameters{1, -1}, &InvalidParametersError{}},
		// Incorrect number of parameters cases
		{Parameters{}, &NumberOfParametersError{}},
		{Parameters{1}, &NumberOfParametersError{}},
		{Parameters{1, 2, 3}, &NumberOfParametersError{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Weibull{%v}", &tt.params)
		t.Run(testname, func(t *testing.T) {
			ans, err := WeibullGenerator(tt.params)

			errType := fmt.Sprintf("%T", err)
			wantErrType := fmt.Sprintf("%T", tt.wantErr)

			if errType != wantErrType {
				t.Errorf("Got %T, wanted %T", err, tt.wantErr)
			}

			if err != nil {
				println(err.Error())
				return
			}

			if ans < 0 {
				t.Errorf("RV should be greater than or equal to 0, but RV = %.2f", ans)
			}

		})

	}

}
