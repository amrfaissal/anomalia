package anomalia

import (
	"math"
)

// Average returns the average of the input
func Average(input []float64) float64 {
	return Sum(input) / float64(len(input))
}

// Sum returns the sum of all elements in the input
func Sum(input []float64) float64 {
	var sum float64
	for _, value := range input {
		sum += value
	}
	return sum
}

// Variance returns the variance of the input
func Variance(input []float64) (variance float64) {
	avg := Average(input)
	for _, value := range input {
		variance += (value - avg) * (value - avg)
	}
	return variance / float64(len(input))
}

// Stdev returns the standard deviation of the input
func Stdev(input []float64) float64 {
	variance := Variance(input)
	return math.Pow(variance, 0.5)
}

// RoundFloat rounds float to closest int
func RoundFloat(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// Float64WithPrecision rounds float to certain precision
func Float64WithPrecision(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(RoundFloat(num*output)) / output
}

// Pdf returns the probability density function
func Pdf(mean, stdev float64) func(float64) float64 {
	return func(x float64) float64 {
		numexp := math.Pow(x-mean, 2) / (2 * math.Pow(stdev, 2))
		denom := stdev * math.Sqrt(2*math.Pi)
		numer := math.Pow(math.E, numexp*-1)
		return numer / denom
	}
}

// Cdf returns the cumulative distribution function
func Cdf(mean, stdev float64) func(float64) float64 {
	return func(x float64) float64 {
		return 0.5 * (1.0 + Erf((x-mean)/(stdev*math.Sqrt(2.0))))
	}
}

// Erf is the guassian error function
func Erf(x float64) float64 {
	// Constants
	a1 := 0.254829592
	a2 := -0.284496736
	a3 := 1.421413741
	a4 := -1.453152027
	a5 := 1.061405429
	p := 0.3275911

	// Save the sign of x
	var sign float64
	if x < 0.0 {
		sign = -1.0
	} else {
		sign = 1.0
	}
	x = math.Abs(x)

	// Formula 7.1.26 given in Abramowitz and Stegun
	t := 1.0 / (1.0 + p*x)
	y := 1.0 - ((((a5*t+a4)*t+a3)*t+a2)*t+a1)*t*math.Pow(math.E, -x*x)
	return sign * y
}