package yfgo

import (
	"fmt"
	"net/url"
)

type QueryParam struct {
	Name  string
	Value string
}

func BuildQuery(params []QueryParam) string {
	queryValues := url.Values{}
	for _, param := range params {
		queryValues.Add(param.Name, param.Value)
	}
	return queryValues.Encode()
}

func MakeURL(url string, symbol string, params []QueryParam) string {
	query := BuildQuery(params)
	return fmt.Sprintf("%s/%s?%s", url, symbol, query)
}
