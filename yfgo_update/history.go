package main

import (
    "fmt"
    "math"
    "reflect"
    "strings"
    "strconv"
    "github.com/jailop/yfgo/yfgo_lib"
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
    content, err := yfgo_lib.ParseJSON(body)
    if err != nil {
        fmt.Println(err)
        return History{}, err
    }
    return History{
        content.Chart.Result[0].TimeStamp,
        NaNZeroPrices(content.Chart.Result[0].Indicators.Quote[0].Open),
        NaNZeroPrices(content.Chart.Result[0].Indicators.Quote[0].Low),
        NaNZeroPrices(content.Chart.Result[0].Indicators.Quote[0].High),
        NaNZeroPrices(content.Chart.Result[0].Indicators.Quote[0].Close),
        content.Chart.Result[0].Indicators.Quote[0].Volume,
    }, nil
}

func GetHistory(symbol string, start_time int64, end_time int64) (History, error) {
    params := []yfgo_lib.QueryParam{
        {Name: "interval", Value: "1m"},
        {Name: "period1", Value: strconv.FormatInt(start_time, 10)},
        {Name: "period2", Value: strconv.FormatInt(end_time, 10)},
    }
    baseURL := "https://query2.finance.yahoo.com/v8/finance/chart"
    url := yfgo_lib.MakeURL(baseURL, symbol, params)
    body, err := yfgo_lib.RetrieveJSON(url)
    if err != nil {
        return History{}, err
    }
    history, err := GenerateHistoryFromParsedJSON(body)
    if err != nil {
        return History{}, err
    }
    return history, nil
}

func NaNZeroPrices(prices []float64) []float64 {
    modPrices := make([]float64, len(prices)) 
    for i := range prices {
        if (prices[i] == 0) {
            modPrices[i] = math.NaN()
        } else {
            modPrices[i] = prices[i]
        }
    }
    return modPrices
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
