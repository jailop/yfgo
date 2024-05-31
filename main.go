package main

import (
    "database/sql"
    "log"
    _ "github.com/marcboeker/go-duckdb"
)

func main() {
    db, err := sql.Open("duckdb", "")
  if err != nil {
    log.Fatal("Failed to connect to database:", err)
  }
  defer db.Close()
}
