package yfgo_lib

import "testing"

func TestUnixTimeAgo(t *testing.T) {
    then, now := UnixTimeAgoAndNow(30 * 60 * 24)
    if then == 0 {
        t.Errorf("unixTimeAgo failed")
    }
    if now == 0 {
        t.Errorf("unixTimeAgo failed")
    }
}
