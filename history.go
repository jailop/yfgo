package main

import (
    "fmt"
    "reflect"
    "strings"
    "strconv"
    _ "database/sql"
    _ "log"
    _ "github.com/marcboeker/go-duckdb"
)

type History struct {
    Time []int64 
    Open []float64
    Low []float64
    High []float64
    Close []float64
    Volume []int64
}

func GenerateHistoryFromParsedJSON(body []byte) (History, error) {
    content, err := ParseJSON(body)
    if err != nil {
        fmt.Println(err)
        return History{}, err
    }
    return History{
        content.Chart.Result[0].TimeStamp,
        content.Chart.Result[0].Indicators.Quote[0].Open,
        content.Chart.Result[0].Indicators.Quote[0].Low,
        content.Chart.Result[0].Indicators.Quote[0].High,
        content.Chart.Result[0].Indicators.Quote[0].Close,
        content.Chart.Result[0].Indicators.Quote[0].Volume,
    }, nil
}

func GetHistory(symbol string, start_time int64, end_time int64) (History, error) {
    params := []QueryParam{
        {name: "interval", value: "1m"},
        {name: "period1", value: strconv.FormatInt(start_time, 10)},
        {name: "period2", value: strconv.FormatInt(end_time, 10)},
    }
    baseURL := "https://query2.finance.yahoo.com/v8/finance/chart"
    url := MakeURL(baseURL, symbol, params)
    println(url)
    body, err := RetrieveJSON(url)
    if err != nil {
        return History{}, err
    }
    history, err := GenerateHistoryFromParsedJSON(body)
    if err != nil {
        return History{}, err
    }
    return history, nil
}

func InsertStatement() string {
    h := new(History)
    query := "INSERT INTO history (symbol,"
    questionMarks := "?,"
    fields := reflect.TypeOf(*h).NumField()
    for i := range fields {
        query += strings.ToLower(reflect.TypeOf(*h).Field(i).Name) + ","
        questionMarks += "?,"
    }
    query = query[:len(query) - 1] + ") VALUES ("
    query += questionMarks[:len(questionMarks) - 1] + ")"
    return query
}
