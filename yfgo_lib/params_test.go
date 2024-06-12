package yfgo_lib

import "testing"

func TestBuildQuery(t *testing.T) {
	tests := []struct {
		Name     string
		params   []QueryParam
		expected string
	}{
		{"empty", nil, ""},
		{"single param", []QueryParam{{Name: "key1", Value: "Value1"}}, "key1=Value1"},
		{"multiple params", []QueryParam{
			{Name: "key1", Value: "Value1"},
			{Name: "key2", Value: "Value2"},
			{Name: "key3", Value: "Value3"},
		}, "key1=Value1&key2=Value2&key3=Value3"},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			query := BuildQuery(test.params)
			if query != test.expected {
				t.Errorf("buildQuery returned an incorrect query: %s", test.expected)
			}
		})
	}
}

func TestMakeURL(t *testing.T) {
	tests := []struct {
		Name     string
		url      string
		symbol   string
		params   []QueryParam
		expected string
	}{
		{
			"empty params",
			"https://example.com",
			"stocks",
			nil,
			"https://example.com/stocks?",
		}, {
			"single param",
			"https://example.com",
			"stocks",
			[]QueryParam{{Name: "symbol", Value: "AAPL"}},
			"https://example.com/stocks?symbol=AAPL",
		}, {
			"multiple params",
			"https://example.com",
			"stocks",
			[]QueryParam{
				{Name: "symbol", Value: "AAPL"},
				{Name: "market", Value: "NASDAQ"},
			},
			"https://example.com/stocks?market=NASDAQ&symbol=AAPL",
		}, {
			"url and symbol only",
			"https://api.example.com",
			"dividends",
			nil,
			"https://api.example.com/dividends?",
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := MakeURL(test.url, test.symbol, test.params)
			if url != test.expected {
				t.Errorf("makeURL returned incorrect URL: %s - %s", url, test.expected)
			}
		})
	}
}
