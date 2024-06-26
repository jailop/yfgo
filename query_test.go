package yfgo

import (
	"math"
	"testing"
)

func TestQuery(t *testing.T) {
	symbol := "IBM"
	start := BackMinutes(7200)
	end := BackMinutes(4320)
	history, err := QueryDB(symbol, start, end)
	if err != nil {
		t.Errorf("Query was not succesful %s", err)
	}
	flag := false
	for i := range len(history.Close) {
		if math.IsNaN(history.Close[i]) {
			flag = true
		}
	}
	if flag {
		t.Errorf("Missed values remaining")
	}
}
