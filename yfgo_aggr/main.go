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
    aggr := history.Aggregate(120)
    aggr.Print()
    println()
    ma := aggr.MovingAverage(9)
    ma.Print()
}
