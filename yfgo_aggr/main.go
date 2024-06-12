package main

import (
    "github.com/jailop/yfgo/yfgo_lib"
)

func main() {
    // history, err := yfgo_lib.QueryDB("BTC-USD", 1710421400, 1717427340)
    history, err := yfgo_lib.QueryDB("BTC-USD", 1710421400, 1718135965)
    if err != nil {
        println(err)
        return
    }
    aggr := AggregateHistory(history, 720)
    aggr.Print()
}
