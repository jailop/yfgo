package main

import (
    "github.com/jailop/yfgo/yfgo_lib"
)

func main() {
    history, err := yfgo_lib.QueryDB("NFLX", 0, yfgo_lib.Now())
    if err != nil {
        println(err)
        return
    }
    // aggr := history.AggregateByDay()
    aggr := history.Aggregate(5)
    // aggr.Print()
    // println()
    dh := aggr.DailyTimedHistory(840)
    dh.Print()
    println()
    ma := dh.MovingAverage(9)
    ma.Print()
    println()
}
