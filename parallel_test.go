package parallel

import (
	"fmt"
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

func TestMap(t *testing.T) {
    data := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

    
    result := Map(data, func(val int) float32 { return float32(val) / 2.0 })
    out := result.Interface().([]float32)
    fmt.Println(out[0]+out[1])

    if out[1] != 1.0 {
        t.Error("Expected data[0]==4, got ", out[1])
    }

}

