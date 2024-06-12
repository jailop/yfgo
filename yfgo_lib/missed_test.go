package yfgo_lib

import (
    "testing"
    "math"
)

func TestFillMissed(t *testing.T) {
    values := [...]float64{5.0, math.NaN(), math.NaN(), 7.0}
    res := FillMissed(values[:], 0.0)
    if res[1] != 6.0 {
        t.Errorf("Filling missed values failed: valid numbers below and above")
    }
    values = [...]float64{math.NaN(), math.NaN(), math.NaN(), 7.0}
    res = FillMissed(values[:], 0.0)
    if res[0] != 7.0 {
        t.Errorf("Filling missed values failed: valid number above %f", res[1])
    }
    values = [...]float64{3.0, math.NaN(), math.NaN(), math.NaN()}
    res = FillMissed(values[:], 0.0)
    if res[3] != 3.0 {
        t.Errorf("Filling missed values failed: valid number above %f", res[3])
    }
    values = [...]float64{math.NaN(), math.NaN(), math.NaN(), math.NaN()}
    res = FillMissed(values[:], 0.0)
    if res[3] != 0.0 {
        t.Errorf("Filling missed values failed: valid number above %f", res[3])
    }
}
