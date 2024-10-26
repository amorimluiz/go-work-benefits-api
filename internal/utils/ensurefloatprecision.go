package utils

import "math"

func EnsureFloatPrecision(value float64, precision int) float64 {
	roundingFactor := math.Pow(10, float64(precision))
	return math.Round(value*roundingFactor) / roundingFactor
}
