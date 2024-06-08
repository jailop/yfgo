package main

import (
    "time"
)

func BackMinutes(minutes int64) int64 {
    return time.Now().Add(-time.Duration(minutes) * time.Minute).Unix()
}

func DefaultThen() int64 {
    const maxPeriod = 7 * 60 * 24
    return BackMinutes(maxPeriod)
}

func Now() int64 {
    return time.Now().Unix()
}

func UnixTimeAgoAndNow(minutes int64) (int64, int64) {
    const maxPeriod int64 = 7 * 60 * 24
    if minutes > maxPeriod {
        minutes = maxPeriod
    }
    now := time.Now()
    then := now.Add(-time.Duration(minutes) * time.Minute)
    return then.Unix(), now.Unix()
}
