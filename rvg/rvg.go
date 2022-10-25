/* Package rvg  provides tools to generate random variates that
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

type RVGenerator func(Parameters) (float64, error)

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

/* Bernoulli Generator takes a univform RV, u, and generates a Bernoulli
   RV.

   The generator function was calculated based on the Bernoulli pmf.
*/
// func BernoulliGenerator(p float64) float64 {
func BernoulliGenerator(params Parameters) (float64, error) {

	if paramsErr := paramsValidate(params, "bernoulli"); paramsErr != nil {
		return 0, paramsErr
	}

	p := params[0]

	u := rand.Float64()

	if u <= 1-p {
		return 0, nil
	} else {
		return 1, nil
	}
}

// func BinomialGenerator(p float64, n int) float64 {
func BinomialGenerator(params Parameters) (float64, error) {

	if paramsErr := paramsValidate(params, "binomial"); paramsErr != nil {
		return 0, paramsErr
	}

	n := int(params[0])
	p := params[1:]

	sum := 0.0
	var trial float64

	for i := 0; i < n; i++ {
		trial, _ = BernoulliGenerator(p)
		sum += trial
	}

	return sum, nil
}

/* ExponentialGenerator takes a uniform RV, u, and generates an eponential
   RV with a given lambda value.

   The generator function was calculated using the inverse transform method
   from the Exponential distrubution's cumulative distribution function.
*/
func ExponentialGenerator(params Parameters) (float64, error) {

	if paramsErr := paramsValidate(params, "exponential"); paramsErr != nil {
		return 0, paramsErr
	}

	λ := params[0]

	u := rand.Float64()

	return -math.Log(1-u) / λ, nil
}

/* See https://www.hongliangjie.com/2012/12/19/how-to-generate-gamma-random-variables/
for implementation */
func GammaGenerator(params Parameters) (float64, error) {

	if paramsErr := paramsValidate(params, "gamma"); paramsErr != nil {
		return 0, paramsErr
	}

	k := params[0]
	θ := params[1]

	if k < 1 {
		gamma, _ := GammaGenerator(Parameters{1.0 + k, θ})
		return gamma * math.Pow(rand.Float64(), 1.0/k), nil
	} else {

		var x, v, u float64
		d := k - 1.0/3.0
		c := 1.0 / math.Sqrt(9*d)

		cont := true

		for cont {
			x = rand.NormFloat64()
			if x > -1/c {
				v = math.Pow(1.0+c*x, 3)

				u = rand.Float64()

				if u < 1-0.0331*x*x*x*x {
					cont = false
				}

				if math.Log(u) < 0.5*x*x+d*(1-v+math.Log(v)) {
					cont = false
				}
			}
		}
		return θ * d * v, nil
	}
}

/* GeometricGenerator takes a probability of success for a Bernoulli trial, p,
   and returns a RV that adheres to the geometric distribution for the given
   value of p.

   The generator function was calculated using the inverse transform method
*/
func GeometricGenerator(params Parameters) (float64, error) {

	if paramsErr := paramsValidate(params, "geometric"); paramsErr != nil {
		return 0, paramsErr
	}

	p := params[0]

	u := rand.Float64()

	return math.Floor(math.Log(u) / math.Log(1-p)), nil
}

/* NormalGenerator returns a pair of normal RVs
   with means mu and deviances sigma^2.

   The generator function is based on the Box-Muller method
*/
// func NormalGenerator(mu, sigma float64) (float64, float64) {
func NormalGenerator(params Parameters) (float64, error) {

	if paramsErr := paramsValidate(params, "normal"); paramsErr != nil {
		return 0, paramsErr
	}

	μ := params[0]
	σ := params[1]

	// u := StdNormalGenerator()
	u := math.Sqrt(-2*math.Log(rand.Float64())) * math.Cos(2*math.Pi*rand.Float64())
	return u*σ + μ, nil
}

/* StdNormalGenerator returns a pair of standard normal RVs.

   The generator function is based on the Box-Muller method
*/
func StdNormalGenerator() float64 {

	u := rand.Float64()
	v := rand.Float64()

	// We can compute two values here, but we only ever use one
	// Return just one
	return math.Sqrt(-2*math.Log(u)) * math.Cos(2*math.Pi*v)
}

/* See http://www.columbia.edu/~ks20/4404-Sigman/4404-Notes-ITM.pdf*/
func PoissonGenerator(params Parameters) (float64, error) {

	if paramsErr := paramsValidate(params, "poisson"); paramsErr != nil {
		return 0, paramsErr
	}

	λ := params[0]

	x := 0.0
	p := rand.Float64()

	for p >= math.Exp(-λ) {
		x++
		p *= rand.Float64()
	}

	return x, nil
}

func WeibullGenerator(params Parameters) (float64, error) {

	if paramsErr := paramsValidate(params, "weibull"); paramsErr != nil {
		return 0, paramsErr
	}

	λ := params[0]
	k := params[1]

	u := rand.Float64()

	return λ * math.Pow(-math.Log(u), 1.0/k), nil
}
