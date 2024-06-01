package main

import (
    "testing"
)

func TestUpdate(t *testing.T) {
    err := Update("NFLX")
    if err != nil {
        t.Errorf("Error updating values %s", err)
    }
}
