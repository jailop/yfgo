package main

import (
    "math"
    "github.com/jailop/yfgo/yfgo_lib"
)

func findOpenPrice(history yfgo_lib.History, start int, end int) float64 {
    res := math.NaN() 
    for i := range (end - start + 1) {
        pos := start + i
        if !math.IsNaN(history.Open[pos]) {
            res = history.Open[pos]
            break
        } else if !math.IsNaN(history.Close[pos]) {
            res = history.Close[pos]
            break
        }
    }
    return res
}

func findClosePrice(history yfgo_lib.History, start int, end int) float64 {
    res := math.NaN()
    for i := (end - start); i >= 0; i-- {
        pos := start + i
        if !math.IsNaN(history.Close[pos]) {
            res = history.Close[pos]
            break
        } else if !math.IsNaN(history.Open[pos]) {
            res = history.Open[pos]
            break
        }
    }
    return res
}

func findHighPrice(history yfgo_lib.History, start int, end int) float64 {
    res := math.Inf(-1)
    for i := range (end - start + 1) {
        pos := start + i
        if !math.IsNaN(history.High[pos]) && history.High[pos] > res {
            res = history.High[pos]
        }
        if !math.IsNaN(history.Open[pos]) && history.Open[pos] > res {
            res = history.Open[pos]
        }
        if !math.IsNaN(history.Close[pos]) && history.Close[pos] > res {
            res = history.Close[pos]
        }
    }
    return res 
}

func findLowPrice(history yfgo_lib.History, start int, end int) float64 {
    res := math.Inf(1)
    for i := range (end - start + 1) {
        pos := start + i
        if !math.IsNaN(history.Low[pos]) && history.Low[pos] < res {
            res = history.High[pos]
        }
        if !math.IsNaN(history.Open[pos]) && history.Open[pos] < res {
            res = history.Open[pos]
        }
        if !math.IsNaN(history.Close[pos]) && history.Close[pos] < res {
            res = history.Close[pos]
        }
    }
    return res 
}

func sumVolume(values []int64) int64 {
    var res int64 = 0
    for i := range len(values) {
        res += values[i]
    }
    return res
}

func aggregatePoint(history yfgo_lib.History, start int, end int) yfgo_lib.HistoryRecord {
    return yfgo_lib.HistoryRecord {
        history.Time[end],
        findOpenPrice(history, start, end),
        findLowPrice(history, start, end),
        findHighPrice(history, start, end),
        findClosePrice(history, start, end),
        sumVolume(history.Volume[start : end + 1]),
    }    
}

// Performs an aggregation on time
func AggregateHistory(history yfgo_lib.History, minutes int64) yfgo_lib.History {
    if (minutes <= 0) {
        return history.Clone()
    } 
    var aggr yfgo_lib.History
    prev := 0
    for i := range len(history.Time) {
        if (history.Time[i] % (minutes * 60) == 0) {
            rec := aggregatePoint(history, prev, i)
            aggr.Append(rec)
            prev = i + 1
        } else if (history.Time[i] - history.Time[prev] > (minutes * 60)) {
            prev = i
        }
    }
    return aggr
}


