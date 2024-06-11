package main

import (
    _ "context"
    "fmt"
    "strings"
    "github.com/jailop/yfgo/yfgo_lib"
)

func LastUpdate(symbol string) int64 {
    db, err := yfgo_lib.OpenDB()
    if err != nil {
        return yfgo_lib.DefaultThen()
    }
    defer db.Close()
    row := db.QueryRow("SELECT MAX(time) FROM history WHERE symbol = ?", symbol)
    var value int64
    err = row.Scan(&value)
    if err != nil {
        fmt.Printf("Error parsing last update for %s\n", symbol)
        return yfgo_lib.DefaultThen()
    }
    if value < yfgo_lib.BackMinutes(7 * 60 * 24) {
        value = yfgo_lib.DefaultThen()
    }
    return value
}

func UpdateTicker(symbol string) error {
    symb := strings.ToUpper(symbol)
    if symb == "" {
        return nil
    }
    db, err := yfgo_lib.OpenDB()
    if err != nil {
        return err
    }
    defer db.Close()
    start_time := LastUpdate(symb) + 1
    end_time := yfgo_lib.Now()
    history, err := GetHistory(symb, start_time, end_time)
    if err != nil {
        return err
    }
    stmt := `
        INSERT INTO 
            history(symbol, time, open, low, high, close, volume)
            VALUES (?, ?, ?, ?, ?, ?, ?)
    `
    counter := 0
    for i := range len(history.Time) {
        if (history.Time[i] < start_time) {
            continue
        } 
        _, err := db.Exec(stmt,
            symb,
            history.Time[i],
            history.Open[i],
            history.Low[i],
            history.High[i],
            history.Close[i],
            history.Volume[i],
        )
        if err != nil {
            fmt.Printf("Record for %d couldn't be inserted\n", history.Time[i])
        } else {
            counter += 1
        }
    }
    fmt.Printf("%s: %d records updated\n", symb, counter)
    return nil
}
