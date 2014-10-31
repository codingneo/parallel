package parallel

import (
    //"fmt"
    "runtime"
    "math"
    //"time"
    //"fmt"
)

type Iterator struct {
    Start int
    End int
    Step int
}

func config() int {
    nCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(nCPU)
    return nCPU   
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
            lb := iter.Start+int(part)*i*iter.Step
            ub := int(math.Min(float64(iter.Start+int(part)*(i+1)*iter.Step),
                               float64(iter.End)))

            for k := lb; k < ub; k = k + iter.Step {
                block(k)
            }

            c <- i
    	}(i)
    }
    
    for i := 0; i < nCPU; i++ {
        //k := <- c
        <- c
        //fmt.Printf("%d part finished ...\n", k)
    }
}

// Parallel For Loop 
func (iter Iterator) Parallelize(block func(int, int) float64) float64 {
    nCPU := config()

    if iter.Step == 0 {
        iter.Step = 1
    }
    var part = int(math.Ceil(float64(iter.End-iter.Start)/float64(iter.Step*nCPU)))
    
    c := make(chan float64, nCPU)
    for i := 0; i < nCPU; i++ {
        go func(i int) {
            lb := iter.Start+part*i*iter.Step
            ub := int(math.Min(float64(iter.Start+part*(i+1)*iter.Step),
                                float64(iter.End)))

            c <- block(lb, ub)
        }(i)
    }
    
    var result float64
    for i := 0; i < nCPU; i++ {
        //k := <- c
        result += <- c
        //fmt.Printf("%d part finished ...\n", k)
    }

    return result
}