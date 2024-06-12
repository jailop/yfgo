package yfgo_lib

import (
    "testing"
)

func TestGetHistory(t *testing.T) {
    timeAgo := DefaultThen()
    then, now := UnixTimeAgoAndNow(timeAgo)
    _, err := GetHistory("COIN", then, now)
    if err != nil {
        t.Errorf("History cannot be retrieved")
    }
}


