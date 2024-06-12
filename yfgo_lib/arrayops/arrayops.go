package arrayops

import (
	"math"
)

func SumArray[T int64 | float64](values []T) T {
	var res T = 0
	for i := range len(values) {
		res += values[i]
	}
	return res
}

// AverageArray computes the averaga of a slice.
// It ignores NaN values.
func AverageArray[T int64 | float64](values []T) float64 {
    var res float64 = 0
    counter := 0
	for i := range len(values) {
        value := float64(values[i])
        if !math.IsNaN(value) {
            res += value
            counter += 1
        }
	}
    return res / float64(counter)
}

func MovingAverage[T int64 | float64](values []T, factor int) []float64 {
	n := len(values)
	res := make([]float64, n)
	for i := range n {
		if i < factor - 1 {
			res[i] = math.NaN()
		} else {
			res[i] = AverageArray(values[i-factor+1:i+1])
		}
	}
	return res
}
