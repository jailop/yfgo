package yfgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type QuoteRec struct {
	Open   []float64
	Low    []float64
	High   []float64
	Close  []float64
	Volume []int64
}

type IndicatorRec struct {
	Quote []QuoteRec `json:"quote"`
}

type ResultRec struct {
	TimeStamp  []int64      `json:"timestamp"`
	Indicators IndicatorRec `json:"indicators"`
}

type ChartRec struct {
	Result []ResultRec `json:"result"`
}

type Content struct {
	Chart ChartRec `json:"chart"`
}

func RetrieveJSON(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	return body, nil
}

func ParseJSON(body []byte) (Content, error) {
	var content Content
	err := json.Unmarshal(body, &content)
	if err != nil {
		fmt.Println(err)
		return content, err
	}
	return content, nil
}

func SaveJSON(body []byte) error {
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.Write(body)
	return err
}
