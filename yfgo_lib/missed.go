package yfgo_lib

import "math"

func computeMissedValue(prev float64, next float64, fail float64) float64 {
    if math.IsNaN(prev) && math.IsNaN(next) {
        return fail
    } else if (math.IsNaN(prev)) {
        return next
    } else if (math.IsNaN(next)) {
        return prev
    } else {
        return (prev + next) / 2
    }
}

func FillMissed(values []float64, fail float64) []float64 {
    n := len(values)
    res := make([]float64, n)
    prevValue := math.NaN()
    for i := range n {
        if math.IsNaN(values[i]) {
            nextValue := math.NaN()
            j := i + 1
            for j < n {
                if !math.IsNaN(values[j]) {
                    nextValue = values[j]
                    break
                }
                j += 1
            }
            res[i] = computeMissedValue(prevValue, nextValue, fail)
        } else {
            res[i] = values[i]
            prevValue = values[i]
        }
    }
    return res
}
