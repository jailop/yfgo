package main

import (
    "log"
    "os"
    "bufio"
    "time"
)

func main() {
    if !DBFileExists() {
        CreateDB()
    }
    listPath, err := FilePath("list.txt")
    if err != nil {
        println(err)
        return
    }
    if !FileExists(listPath) {
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
