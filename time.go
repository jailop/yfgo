package main

import (
    "time"
)

func UnixTimeAgoAndNow(minutes int) (int64, int64) {
    maxPeriod := 7 * 60 * 24
    if minutes > maxPeriod {
        minutes = maxPeriod
    }
    now := time.Now()
    then := now.Add(-time.Duration(minutes) * time.Minute)
    return then.Unix(), now.Unix()
}
