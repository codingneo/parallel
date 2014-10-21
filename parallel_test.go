package parallel

import (
	"fmt"
	"testing"
    "time"
)

func TestFor(t *testing.T) {
    data := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

    For(Iterator{0, 15, 1},
        func(i int) { data[i] = data[i]*2 })

    if data[0] != 2 {
        t.Error("Expected data[0]==2, got ", data[0])
    }
}

func TestPerformance(t *testing.T) {
    len := 100000000
    data := make([]int, len)

    t0 := time.Now()

    for i := 0; i < len; i++ {
        data[i] = 3.0
    }

    t1 := time.Now()
    fmt.Printf("The call1 took %v to run.\n", t1.Sub(t0))

    t0 = time.Now()
    For(Iterator{Start: 0, End: len},
        func(i int) { data[i] = 4.0 })
    t1 = time.Now()
    fmt.Printf("The call2 took %v to run.\n", t1.Sub(t0))


    if (data[3] != 4.0) {
        t.Error("Expected data[3]==4, got ", data[3])
    }
}

