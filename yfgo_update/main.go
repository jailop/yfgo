package main

import (
    "log"
    "os"
    "bufio"
    "time"
    "github.com/jailop/yfgo/yfgo_lib"
)

func main() {
    if !yfgo_lib.DBFileExists() {
        yfgo_lib.CreateDB()
    }
    listPath, err := yfgo_lib.FilePath("list.txt")
    if err != nil {
        println(err)
        return
    }
    if !yfgo_lib.FileExists(listPath) {
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
        UpdateTicker(scanner.Text())
        time.Sleep(3 * time.Second)
    }
    if err = scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
