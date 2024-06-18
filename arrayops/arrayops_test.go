package arrayops

import (
	"math"
	"testing"
)

func TestMovingAverage(t *testing.T) {
	a := [...]float64{1.0, 2.0, 3.0, 4.0, 5.0}
	b := MovingAverage(a[:], 3)
	if !math.IsNaN(b[0]) {
		t.Error("Non generating nan for non complete series")
	}
    if !math.IsNaN(b[1]) {
        t.Errorf("Expected for position 1: NaN - %f\n", b[1])
    }
	if b[2] != 2.0 {
		t.Errorf("Bad computation on moving average: %f - %f", 2.0, b[2])
	}
	if b[3] != 3.0 {
		t.Errorf("Bad computation on moving average: %f - %f", 3.0, b[3])
	}
	if b[4] != 4.0 {
		t.Errorf("Bad computation on moving average: %f - %f", 4.0, b[4])
	}
}
