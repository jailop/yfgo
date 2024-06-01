# Yahoo Finance Minute Data Retriever

Retrieves minute stock prices from Yahoo Finance.

Data is saved in a DuckDB file under the directory `$HOME/.config/yfgo`

In that folder, the user should create a file named `list.txt` that lists
ticker symbols for which data will be retrieved. Write a symbol per line.

To build and install the binary:

```bash
go build
go install
```

Make sure to add the `$HOME/go/bin` to your path.

This program can be schedule to run periodically using a cron tab, like the
below one. That indicates that the app will run every 5 minutes from 9 to 16
hours between Monday and Friday.

```cron
*/5 9-16 * * MON-FRI $HOME/bin/yfgo
```

Only needed data will be retrieved from Yahoo Finance.

To access the database, you can use any DuckDB interface. For example:

```bash
duckdb $HOME/.config/yfgo/data.db
```

The database only contains a one table: `history`:

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

* Yahoo Finance are registered trademarks of Yahoo, Inc.
* The author and this project is not affiliated, endorsed, or vetted by Yahoo Inc.
* This app uses Yahoo's publicly available API's.
* This app is intended for research and educational purposes.
