package arrayops

import (
    "sort"
)

type Pair = struct {
    idx uint
    value int64
}

func SortedIndices(data []int64) []uint {
    pairs := make([]Pair, len(data))
    pos := 0
    for i := range data {
        pairs[i] = Pair{uint(i), data[i]}
        pos += 1
    }
    sort.Slice(pairs, func (i, j int) bool {
        return pairs[i].value < pairs[j].value
    })
    indices := make([]uint, len(data))
    for i := range pairs {
        indices[i] = pairs[i].idx
    }
    return indices
}
