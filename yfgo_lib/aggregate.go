package yfgo_lib

import (
	"math"
    "github.com/jailop/yfgo/yfgo_lib/arrayops"
)

const day = 24 * 60 * 60

func (history History) findOpenPrice(start int, end int) float64 {
	res := math.NaN()
	for i := range end - start + 1 {
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

func (history History) findClosePrice(start int, end int) float64 {
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

func (history History) findHighPrice(start int, end int) float64 {
	res := math.Inf(-1)
	for i := range end - start + 1 {
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

func (history History) findLowPrice(start int, end int) float64 {
	res := math.Inf(1)
	for i := range end - start + 1 {
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

func (history History) aggregatePoint(start int, end int) HistoryRecord {
	return HistoryRecord{
		history.Time[end],
		history.findOpenPrice(start, end),
		history.findLowPrice(start, end),
		history.findHighPrice(start, end),
		history.findClosePrice(start, end),
		arrayops.SumArray(history.Volume[start : end+1]),
	}
}

// Aggregate performs an aggregation on time.
// It generates price and volume values on intervals of time of length `minutes`.
func (history History) Aggregate(minutes int64) History {
	if minutes <= 1 {
		return history.Clone()
	}
	var aggr History
	prev := 0
	for i := range len(history.Time) {
		if history.Time[i]%(minutes*60) == 0 {
			rec := history.aggregatePoint(prev, i)
			aggr.Append(rec)
			prev = i + 1
		} else if history.Time[i]-history.Time[prev] > (minutes * 60) {
			prev = i
		}
	}
	return aggr
}

func floorDay(when int64) int64 {
	return when / day * day
}

func ceilDay(when int64) int64 {
	return when/day*day + day - 1
}

// AggregateByDay performs an aggregation by day.
// It generates price and volume values for every day in the given interval.
func (history History) AggregateByDay() History {
	var aggr History
	if len(history.Time) == 0 {
		return History{}
	}
	mark := floorDay(history.Time[0]) + day
	prev := 0
	i := 0
	for i < len(history.Time) {
		if history.Time[i] >= mark {
			rec := history.aggregatePoint(prev, i-1)
			rec.Time = mark - 60
			aggr.Append(rec)
			mark += day
			prev = i
		}
		i += 1
	}
	if prev != i-1 {
		rec := history.aggregatePoint(prev, len(history.Time)-1)
		rec.Time = mark - 60
		aggr.Append(rec)
	}
	return aggr
}

func (history History) MovingAverage(factor int) History{
    if factor > len(history.Time) {
        return History{}
    }
    pos := factor - 1
    Open := arrayops.MovingAverage(history.Open, factor)[pos:]
    Low := arrayops.MovingAverage(history.Low, factor)[pos:]
    High := arrayops.MovingAverage(history.High, factor)[pos:]
    Close := arrayops.MovingAverage(history.Close, factor)[pos:]
    Volume := arrayops.MovingAverage(history.Volume, factor)[pos:]
    VolumeInt := make([]int64, len(Volume))
    for i := range len(Volume) {
        VolumeInt[i] = int64(Volume[i])
    }
    return History{
        history.Time[pos:],
        Open,
        Low,
        High,
        Close,
        VolumeInt,    
    }
}
