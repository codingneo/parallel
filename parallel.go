package parallel

import (
    "fmt"
)

type Iterator struct {
    Start int
    End int
    Step int
}

func For(iter Iterator, block func(int)) {
    factor := 4
    if iter.Step == 0 {
        iter.Step = 1
    }
    var part = float32((iter.End-iter.Start)/(iter.Step*factor))
    
    c := make(chan int)
    for i := 0; i < factor; i++ {
    	go func(i int) {
            for k := iter.Start+int(part)*i*iter.Step; 
                k < int(part)*(i+1)*iter.Step; 
                k = k + iter.Step {
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

