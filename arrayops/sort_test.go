package arrayops

import "testing"

func TestSortedIndices(t *testing.T) {
    a := [...]int64{6, 3, 4}
    a_ := a[0:]
    b := SortedIndices(a_)
    if b[0] != 1 {
        t.Error("Bad classification")
    }
    if b[1] != 2 {
        t.Error("Bad classification")
    }
    if b[2] != 0 {
        t.Error("Bad classification")
    }
}
