package parallel

import (
    "fmt"
    "runtime"
    "math"
)

type Iterator struct {
    Start int
    End int
    Step int
}

func config() int {
    runtime.GOMAXPROCS(runtime.NumCPU())
    return runtime.NumCPU()    
}

// Parallel For Loop 
func (iter Iterator) For(block func(int)) {
    nCPU := config()

    if iter.Step == 0 {
        iter.Step = 1
    }
    var part = math.Ceil(float64(iter.End-iter.Start)/float64(iter.Step*nCPU))
    
    c := make(chan int, nCPU)
    for i := 0; i < nCPU; i++ {
    	go func(i int) {
            ubound := math.Min(float64(int(part)*(i+1)*iter.Step),
                                float64(iter.End))
            for k := iter.Start+int(part)*i*iter.Step; 
                k < int(ubound); 
                k = k + iter.Step {
                block(k)
            }
            c <- i
    	}(i)
    }
    
    for i := 0; i < nCPU; i++ {
        k := <- c
        fmt.Printf("%d part finished ...\n", k)
    }
}