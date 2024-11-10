/*
This program shows how to use the yfgo's aggregation functionality.
By default, data retrieve by yfgo is by minute and it is saved in a
DuckDB database. At the same time, yfgo provides the function
`Aggregate` that can be applied over an history object. It takes as
argument the number of minutes representing the intervals of time
on which you'll get the data.
*/
package main

import (
    "github.com/jailop/yfgo"
)

func main() {
	// Retrieving the full history available for a symbol
    history, err := yfgo.QueryDB("BTC-USD", 0, yfgo.Now())
    if err != nil {
        println(err)
        return
    }
	// Generating aggregations
    aggr := history.Aggregate(60)
    aggr.Print()
    println("------------")
	// Printing moving avegerages every hour 
	// based on the previous 9 days an the same time
    dtma := history.DailyTimedMovingAverages(60, 9)
    dtma.Print()
}
