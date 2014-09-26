package parallel

import (
    "fmt"
    "time"
    "math/rand"
)

func For(start, end int, block func(int)) {
    factor := 4
    var part = float32((end-start)/factor)
    
    c := make(chan int)
    for i := 0; i < factor; i++ {
    	go func(i int) {
            for k := start+int(part)*i; k < int(part)*(i+1); k++ {
            	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
                block(k)
            }
            c <- i
    	}(i)
    }
    
    for i := 0; i < factor; i++ {
        k := <- c
        fmt.Printf("%d part finished ...\n", k)
    }
}