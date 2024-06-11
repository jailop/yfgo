Retrieve(symbol string, start int64, end int64) {

}

CoveredPeriod(symbol) (int64, int64) {
    db, err := OpenDB()
    if err != nil {
        return DefaultThen()
    }
    defer db.Close()
}
