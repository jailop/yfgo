package main

import (
    "log"
    "os"
    "bufio"
    "time"
    "github.com/jailop/yfgo/dbconn"
    "github.com/jailop/yfgo/fileutils"
)

func main() {
    if !dbconn.DBFileExists() {
        dbconn.CreateDB()
    }
    listPath, err := fileutils.FilePath("list.txt")
    if err != nil {
        println(err)
        return
    }
    if !fileutils.FileExists(listPath) {
        println("List of ticker symbols doesn't exist")
        println("Create a new one at ", listPath)
        return
    }
    file, err := os.Open(listPath)
    if err != nil {
        log.Fatal(err) 
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        Update(scanner.Text())
        time.Sleep(3 * time.Second)
    }
    if err = scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
