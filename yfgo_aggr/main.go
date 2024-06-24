package main

import (
    "github.com/jailop/yfgo"
)

func main() {
    history, err := yfgo.QueryDB("BTC-USD", 0, yfgo.Now())
    if err != nil {
        println(err)
        return
    }
    aggr := history.Aggregate(60)
    aggr.Print()
    println("------------")
    dtma := history.DailyTimedMovingAverages(60, 9)
    dtma.Print()
}
