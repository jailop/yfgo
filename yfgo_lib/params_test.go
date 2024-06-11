package yfgo_lib

import "testing"

func TestBuildQuery(t *testing.T) {
	tests := []struct {
		name     string
		params   []QueryParam
		expected string
	}{
		{"empty", nil, ""},
		{"single param", []QueryParam{{name: "key1", value: "value1"}}, "key1=value1"},
		{"multiple params", []QueryParam{
			{name: "key1", value: "value1"},
			{name: "key2", value: "value2"},
			{name: "key3", value: "value3"},
		}, "key1=value1&key2=value2&key3=value3"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			query := BuildQuery(test.params)
			if query != test.expected {
				t.Errorf("buildQuery returned an incorrect query: %s", test.expected)
			}
		})
	}
}

func TestMakeURL(t *testing.T) {
	tests := []struct {
		name     string
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
			[]QueryParam{{name: "symbol", value: "AAPL"}},
			"https://example.com/stocks?symbol=AAPL",
		}, {
			"multiple params",
			"https://example.com",
			"stocks",
			[]QueryParam{
				{name: "symbol", value: "AAPL"},
				{name: "market", value: "NASDAQ"},
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
		t.Run(test.name, func(t *testing.T) {
			url := MakeURL(test.url, test.symbol, test.params)
			if url != test.expected {
				t.Errorf("makeURL returned incorrect URL: %s - %s", url, test.expected)
			}
		})
	}
}
