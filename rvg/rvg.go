/* Package rvg  provides tools to generate random variables that
 * fit the following distrubtions: Bernoulli, Geometric, Exponential
 * Normal, Gamma, Weibull
 */
package rvg

import (
	"math"
	"math/rand"
	"time"
)

// InitRNG seeds the math/rand PRNG using the current time.
func InitRNG() {
	rand.Seed(time.Now().UnixMicro())
}

// InitRNGWithSeed seeds the math/rand PRNG using a value passed in
// for repoducability of results.
func InitRNGWithSeed(seed int64) {
	rand.Seed(seed)
}

// :TODO: Replace all generators with dist([]float64) float64 for parameters

type RVGenerator func([]float64) float64

var Generators = map[string]RVGenerator{
	"bernoulli":   BernoulliGenerator,
	"binomial":    BinomialGenerator,
	"exponential": ExponentialGenerator,
	"gamma":       GammaGenerator,
	"geometric":   GeometricGenerator,
	"normal":      NormalGenerator,
	"poisson":     PoissonGenerator,
	"weibull":     WeibullGenerator,
}

/* ExponentialGenerator takes a uniform RV, u, and generates an eponential
   RV with a given lambda value.

   The generator function was calculated using the inverse transform method
   from the Exponential distrubution's cumulative distribution function.
*/
func ExponentialGenerator(params []float64) float64 {

	// :TODO: validate params
	lambda := params[0]

	u := rand.Float64()

	return -math.Log(1-u) / lambda
}

/* Bernoulli Generator takes a univform RV, u, and generates a Bernoulli
   RV.

   The generator function was calculated based on the Bernoulli pmf.
*/
// func BernoulliGenerator(p float64) float64 {
func BernoulliGenerator(params []float64) float64 {

	// :TODO: validate params
	p := params[0]

	u := rand.Float64()

	if u <= 1-p {
		return 0
	} else {
		return 1
	}
}

// func BinomialGenerator(p float64, n int) float64 {
func BinomialGenerator(params []float64) float64 {

	// :TODO: validate params
	p := params[0:1]
	n := int(params[1])

	sum := 0.0

	for i := 0; i < n; i++ {
		sum += BernoulliGenerator(p)
	}

	return sum
}

/* See http://www.columbia.edu/~ks20/4404-Sigman/4404-Notes-ITM.pdf*/
func PoissonGenerator(params []float64) float64 {

	// :TODO: validate parameters
	lambda := params[0]

	X := 0.0
	P := rand.Float64()

	for P >= math.Exp(-lambda) {
		X++
		P *= rand.Float64()
	}

	return X
}

/* StdNormalGenerator returns a pair of standard normal RVs.

   The generator function is based on the Box-Muller method
*/
func StdNormalGenerator() (float64, float64) {

	u := rand.Float64()
	v := rand.Float64()

	return math.Sqrt(-2*math.Log(u)) * math.Cos(2*math.Pi*v), math.Sqrt(-2*math.Log(u)) * math.Sin(2*math.Pi*v)
}

/* NormalGenerator returns a pair of normal RVs
   with means mu and deviances sigma^2.

   The generator function is based on the Box-Muller method
*/
// func NormalGenerator(mu, sigma float64) (float64, float64) {
func NormalGenerator(params []float64) float64 {

	// :TODO: validate parameters
	// if params is empty, give std normal?

	mu := params[0]
	sigma := params[1] // :TODO: Should this be sigma^2?

	u, _ := StdNormalGenerator()

	return u*sigma + mu // :TODO: Should this use sqrt(sigma^2)?
	// return u*sigma + mu, v*sigma + mu
}

/* GeometricGenerator takes a probability of success for a Bernoulli trial, p,
   and returns a RV that adheres to the geometric distribution for the given
   value of p.

   The generator function was calculated using the inverse transform method
*/
// func GeometricGenerator(p float64) float64 {
func GeometricGenerator(params []float64) float64 {

	// :TODO: validate params
	p := params[0]

	u := rand.Float64()

	return math.Log(u) / math.Log(1-p)
}

func GammaGenerator(params []float64) float64 {

	// See https://www.hongliangjie.com/2012/12/19/how-to-generate-gamma-random-variables/
	// for implementation

	// :TODO: validate params
	a := params[0]
	b := params[1]

	if a > 1 {
		return GammaGenerator([]float64{1.0 + a, b}) * math.Pow(rand.Float64(), 1.0/a)
	} else {
		var x, v, u float64
		d := a - 1.0/3.0
		c := (1.0 / 3.0) / math.Sqrt(d)

		for {
			v := 1.0 + c*rand.NormFloat64()
			for v >= 0 {
				v = 1.0 + c*rand.NormFloat64()
			}

			v = v * v * v
			u = rand.Float64()

			if u > 1-0.0331*x*x*x*x {
				break
			}

			if math.Log(u) > 0.5*x*x+d*(1-v+math.Log(v)) {
				break
			}
		}

		return b * d * v

	}
}

func WeibullGenerator(params []float64) float64 {

	// :TODO: validate params
	a := params[0]
	b := params[1]

	u := rand.Float64()

	return math.Pow(-math.Log(u), 1.0/b) / a
}
