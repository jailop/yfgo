package yfgo_lib

import (
    "fmt"
    "log"
)

func QueryDB(symbol string, start int64, end int64) (History, error) {
    db, err := OpenDB()
    if err != nil {
        return History{}, err 
    }
    defer db.Close();
    stmt := `
        SELECT time, open, low, high, close, volume
        FROM history
        WHERE symbol=? AND time>=? AND time<=?
        ORDER BY time
    `
    row, err := db.Query(stmt, symbol, start, end)
    if err != nil {
        fmt.Printf("Data cannot be retrieved for %s\n", symbol)
        fmt.Printf("%s", err)
        return History{}, err
    }
    defer row.Close()
    var history History
    for row.Next() {
        var rec HistoryRecord
        err := row.Scan(&rec.Time, &rec.Open, &rec.Low, &rec.High, &rec.Close, &rec.Volume)
        if err != nil {
            log.Println(err)
            continue
        } 
        history.Append(rec)
    }
    // history.FillMissed()
    return history, nil
}
