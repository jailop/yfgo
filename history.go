package yfgo

import (
	_ "database/sql"
	"fmt"
	_ "log"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jailop/yfgo/arrayops"
	_ "github.com/marcboeker/go-duckdb"
)

// Represent an Stock Price/Volume history
type History struct {
	Time   []int64      `json:"time"`
	Open   []float64    `json:"open"`
	Low    []float64    `json:"low"`
	High   []float64    `json:"high"`
	Close  []float64    `json:"close"`
	Volume []int64      `json:"volume"`
}

type HistoryRecord struct {
	Time   int64
	Open   float64
	Low    float64
	High   float64
	Close  float64
	Volume int64
}

func (history *History) Append(record HistoryRecord) {
	history.Time = append(history.Time, record.Time)
	history.Open = append(history.Open, record.Open)
	history.Low = append(history.Low, record.Low)
	history.High = append(history.High, record.High)
	history.Close = append(history.Close, record.Close)
	history.Volume = append(history.Volume, record.Volume)
}

func (history *History) AppendFromHistory(another History, idx int) {
    history.Time = append(history.Time, another.Time[idx])
    history.Open = append(history.Open, another.Open[idx])
    history.Low = append(history.Low, another.Low[idx])
    history.High = append(history.High, another.High[idx])
    history.Close = append(history.Close, another.Close[idx])
    history.Volume = append(history.Volume, another.Volume[idx])
}

func (history *History) AppendHistory(another History) {
    for idx := range len(another.Time) {
        history.AppendFromHistory(another, idx)
    }
}

func (history *History) FillMissed() {
	history.Open = FillMissed(history.Open, 0.0)
	history.Low = FillMissed(history.Low, 0.0)
	history.High = FillMissed(history.High, 0.0)
	history.Close = FillMissed(history.Close, 0.0)
}

func (history History) Print() {
	for i := range len(history.Time) {
		fmt.Printf("%14s %10.2f %10.2f %10.2f %10.2f %10d\n",
			time.Unix(history.Time[i], 0).UTC(),
			history.Open[i],
			history.Low[i],
			history.High[i],
			history.Close[i],
			history.Volume[i],
		)
	}
}

func copyArray[T int64 | float64](values []T) []T {
	ret := make([]T, len(values))
	for i := range len(values) {
		ret[i] = values[i]
	}
	return ret
}

func (history History) Clone() History {
	return History{
		copyArray(history.Time),
		copyArray(history.Open),
		copyArray(history.Low),
		copyArray(history.High),
		copyArray(history.Close),
		copyArray(history.Volume),
	}
}

func GenerateHistoryFromParsedJSON(body []byte) (History, error) {
	content, err := ParseJSON(body)
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
	params := []QueryParam{
		{Name: "interval", Value: "1m"},
		{Name: "period1", Value: strconv.FormatInt(start_time, 10)},
		{Name: "period2", Value: strconv.FormatInt(end_time, 10)},
	}
	baseURL := "https://query2.finance.yahoo.com/v8/finance/chart"
	url := MakeURL(baseURL, symbol, params)
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

func (history History) Sort() History {
    n := len(history.Time)
    time_ := make([]int64, n)
    open_ := make([]float64, n)
    low_ := make([]float64, n)
    high_ := make([]float64, n)
    close_ := make([]float64, n)
    volume_ := make([]int64, n)
    indices := arrayops.SortedIndices(history.Time)
    for i, pos := range indices {
        time_[i] = history.Time[pos]
        open_[i] = history.Open[pos]
        low_[i] = history.Low[pos]
        high_[i] = history.High[pos]
        close_[i] = history.Close[pos]
        volume_[i] = history.Volume[pos]
    }
    return History{time_, open_, low_, high_, close_, volume_}
}

func (history History) ToJSON() ([]byte, error) {
    return json.Marshal(history, "", "    ")
}

func NaNZeroPrices(prices []float64) []float64 {
	modPrices := make([]float64, len(prices))
	for i := range prices {
		if prices[i] == 0 {
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
	query = query[:len(query)-1] + ") VALUES ("
	query += questionMarks[:len(questionMarks)-1] + ")"
	return query
}
