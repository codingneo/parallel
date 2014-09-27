package parallel

import (
    "testing"
)

func TestFor(t *testing.T) {
    data := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

    For(Iterator{0, 15, 1},
        func(i int) { data[i] = data[i]*2 })

    if data[0] != 2 {
        t.Error("Expected data[0]==2, got ", data[0])
    }
}

