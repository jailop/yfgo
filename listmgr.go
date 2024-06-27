package yfgo

import (
    "log"
    "os"
    "bufio"
    "strings"
)

func SymbolList() []string {
    symlist := make([]string, 0)
    listPath, err := FilePath("list.txt")
    if err != nil {
        log.Fatal(err)
        return symlist
    }
    if !FileExists(listPath) {
        log.Fatal("List of ticker symbols doesn't exist")
        log.Print("Create a new one at ", listPath)
        return symlist
    }
    file, err := os.Open(listPath)
    if err != nil {
        log.Fatal(err) 
        return symlist
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        vs := strings.Split(scanner.Text(), " ")
        for _, sym := range vs {
            symlist = append(symlist, strings.TrimSpace(sym))
        }
    }
    if err = scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return symlist
}

func SymbolExists(sym string) bool {
    symlist := SymbolList()
    flag := false
    for _, val := range symlist {
        if sym == strings.TrimSpace(val) {
            flag = true
            break
        }
    }
    return flag
}

func SymbolAdd(sym string) {
    if SymbolExists(sym) {
        log.Print("SymboAdd: symbol already exists")
        return
    }
    listPath, err := FilePath("list.txt")
    file, err := os.OpenFile(listPath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal("SymbolAdd: File cannot be openned")
        return
    }
    defer file.Close()
    if _, err = file.WriteString(strings.TrimSpace(sym) + "\n"); err != nil {
        log.Fatal("SymbolAdd: Symbol cannot be appended")
        log.Fatal(err)
    }
}

func SymbolListWrite(symlist []string) {
    listPath, err := FilePath("list.txt")
    file, err := os.OpenFile(listPath, os.O_RDWR | os.O_CREATE, 0644)
    if err != nil {
        log.Fatal("SymbolAdd: File cannot be openned")
        return
    }
    for _, val := range symlist {
        if _, err = file.WriteString(strings.TrimSpace(val) + "\n"); err != nil {
            log.Fatal("SymbolAdd: Symbol cannot be appended")
        }
    }
    println("Here >>>")
    if err = file.Close(); err != nil {
        log.Fatal(err)
    }
}

func SymbolRemove(sym string) bool {
    flag := false
    symlist := SymbolList()
    newlist := make([]string, 0)
    for _, val := range symlist {
        if strings.TrimSpace(val) == strings.TrimSpace(sym) {
            flag = true
        } else {
            newlist = append(newlist, val)
        }
    }
    SymbolListWrite(newlist)
    return flag
}
