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

/* ExponentialGenerator takes a uniform RV, u, and generates an eponential
   RV with a given lambda value.

   The generator function was calculated using the inverse transform method
   from the Exponential distrubution's cumulative distribution function.
*/
func ExponentialGenerator(lambda float64) float64 {

	u := rand.Float64()

	return -math.Log(1-u) / lambda
}

/* Bernoulli Generator takes a univform RV, u, and generates a Bernoulli
   RV.

   The generator function was calculated based on the Bernoulli pmf.
*/
func BernoulliGenerator(p float64) float64 {

	u := rand.Float64()

	if u < p {
		return 0
	} else {
		return 1
	}
}

func BinomialGenerator(p float64, n int) float64 {

	sum := 0.0

	for i := 0; i < n; i++ {
		sum += BernoulliGenerator(p)
	}

	return sum
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
func NormalGenerator(mu, sigma float64) (float64, float64) {

	u, v := StdNormalGenerator()

	return u*sigma + mu, v*sigma + mu
}

/* GeometricGenerator takes a probability of success for a Bernoulli trial, p,
   and returns a RV that adheres to the geometric distribution for the given
   value of p.

   The generator function was calculated using the inverse transform method
*/
func GeometricGenerator(p float64) float64 {

	u := rand.Float64()

	return math.Log(u) / math.Log(1-p)
}

func GammaGenerator() float64 {

	// See https://www.hongliangjie.com/2012/12/19/how-to-generate-gamma-random-variables/
	// for implementation

	return 0
}

func WeibullGenerator(a, b float64) float64 {

	u := rand.Float64()

	return math.Pow(-math.Log(u), 1.0/b) / a
}
