# yfgo: Yahoo Finance Data Retriever

Retrieves minute stock prices from Yahoo Finance. This project is inspired by [yfinance](https://github.com/ranaroussi/yfinance).

Data is saved in a DuckDB file under the directory `$HOME/.config/yfgo`

In that folder, the user should create a file named `list.txt` that lists
ticker symbols for which data will be retrieved. Write a symbol per line.

To build and install the binary:

```bash
cd yfgo/yfgo_update
go build
go install
```

Make sure to add the `$HOME/go/bin` to your path.

This program can be scheduled to run periodically using a cron tab, like the
below one. That indicates that the app will run every hour between Monday and Friday.

```cron
* */1 * * MON-FRI $HOME/go/bin/yfgo_update
```

Only needed data will be retrieved from Yahoo Finance.

To access the database, you can use any DuckDB interface. For example:

```bash
duckdb $HOME/.config/yfgo/data.db
```

The database contains only one table: `history`:

```SQL
SELECT * FROM history LIMIT 10;
```

```
┌─────────┬────────────┬───┬────────────────────┬────────────────────┬────────┐
│ symbol  │    time    │ … │        high        │       close        │ volume │
│ varchar │   int64    │   │       double       │       double       │ int64  │
├─────────┼────────────┼───┼────────────────────┼────────────────────┼────────┤
│ COIN    │ 1716903000 │ … │ 235.91000366210938 │  235.5850067138672 │ 415371 │
│ COIN    │ 1716903060 │ … │ 238.13999938964844 │ 238.02000427246094 │ 179997 │
│ COIN    │ 1716903120 │ … │ 239.78970336914062 │ 239.70289611816406 │ 178277 │
│ COIN    │ 1716903180 │ … │ 241.25999450683594 │  240.4600067138672 │ 267419 │
│ COIN    │ 1716903240 │ … │ 241.16000366210938 │  237.0800018310547 │ 179495 │
│ COIN    │ 1716903300 │ … │  238.6999969482422 │ 235.58999633789062 │ 145028 │
│ COIN    │ 1716903360 │ … │ 235.92990112304688 │ 231.74000549316406 │ 162318 │
│ COIN    │ 1716903420 │ … │ 233.86500549316406 │ 233.00010681152344 │ 137697 │
│ COIN    │ 1716903480 │ … │                0.0 │                0.0 │      0 │
│ COIN    │ 1716903540 │ … │              233.5 │  233.1649932861328 │ 109542 │
├─────────┴────────────┴───┴────────────────────┴────────────────────┴────────┤
│ 10 rows                                                 7 columns (5 shown) │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Disclaimers**:

* Yahoo Finance is a registered trademark of Yahoo, Inc.
* The author and this project are not affiliated, endorsed, or vetted by Yahoo Inc.
* This app uses Yahoo's publicly available API's.
* This app is intended for research and educational purposes.
