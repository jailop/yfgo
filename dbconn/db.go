package dbconn

import (
    "database/sql"
    _ "github.com/marcboeker/go-duckdb"
    "github.com/jailop/yfgo/fileutils"
)

func DBPath() (string, error){
    return fileutils.FilePath("data.db")
}

func DBFileExists() bool {
    dbPath, err := DBPath()
    if err != nil {
        return false
    }
    return fileutils.FileExists(dbPath)
}

func OpenDB() (*sql.DB, error) {
    dbFile, err := DBPath()
    if err != nil {
        return nil, err
    }
    db, err := sql.Open("duckdb", dbFile)
    if err != nil {
        println("Database cannot be openned")
        return nil, err
    }
    return db, nil
}

func CreateDB() error {
    db, err := OpenDB()
    if err != nil {
        return err
    }
    defer db.Close()
    _, err = db.Exec("DROP TABLE IF EXISTS history")
    if err != nil {
        println("Old history table cannot be deleted")
        return err
    }
    stmt := `
        CREATE TABLE history (
            symbol VARCHAR,
            time BIGINT,
            open DOUBLE,
            low DOUBLE,
            high DOUBLE,
            close DOUBLE,
            volume BIGINT,
            PRIMARY KEY (symbol, time)
        )
    `
    _, err = db.Exec(stmt)
    if err != nil {
        println("Table history cannot be created")
        return err
    }
    println("Database has been created")
    return nil
}
