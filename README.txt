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

Only need data will be retrieved from Yahoo Finance.
