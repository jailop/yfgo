package main

import (
    "github.com/jailop/yfgo"
)

func main() {
    history, err := yfgo.QueryDB("NFLX", 0, yfgo.Now())
    if err != nil {
        println(err)
        return
    }
    aggr := history.DailyTimedMovingAverages(60, 9)
    aggr.Print()
}
