package main

import (
    "testing"
    // "strconv"
    "io/ioutil"
    "encoding/json"
)

func TestRetrieveJSON(t *testing.T) {
//    timeAgo := 60 * 24 * 7
//    then, now := UnixTimeAgoAndNow(timeAgo)
//    params := []QueryParam{
//        {name: "interval", value: "1m"},
//        {name: "period1", value: strconv.FormatInt(then, 10)},
//        {name: "period2", value: strconv.FormatInt(now, 10)},
//    }
//    baseURL := "https://query2.finance.yahoo.com/v8/finance/chart"
//    url := MakeURL(baseURL, "COIN", params)
//    body, err := RetrieveJSON(url)
//    if err != nil {
//        t.Errorf("Retrieving JSON file failed")
//    }
//    err = SaveJSON(body)
//    if err != nil {
//        t.Errorf("JSON file cannot be saved")
//    }
    body, err := ioutil.ReadFile("output.json")
    if err != nil {
        t.Errorf("JSON file cannot be read")
        return
    }
    content, err := ParseJSON(body)
    if err != nil {
        t.Errorf("JSON cannot be parsed")
        return
    }
    _, err = json.Marshal(content)
    if err != nil {
        t.Errorf("JSON cannot be restored")
        return
    }
    history, err := GenerateHistoryFromParsedJSON(body)
    if err != nil {
        t.Errorf("History cannot be generated")
        return
    }
    if len(history.Open) == 0 {
        t.Errorf("Open does not have any data value")
        return
    }
    println(InsertStatement())
}
