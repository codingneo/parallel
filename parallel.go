package parallel

import (
    "fmt"
    "reflect"
)

type Iterator struct {
    Start int
    End int
    Step int
}

//function types
type mapf func(interface{}) interface{}
type reducef func(interface{}, interface{}) interface{}
type filterf func(interface{}) bool

// Parallel For Loop 
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


func Map(in interface{}, fn mapf) reflect.Value {
    // TODO: dynamic choose parallel factor
    factor := 4

    val := reflect.ValueOf(in)
    ftype := reflect.ValueOf(fn).Type()
    out := reflect.MakeSlice(
            reflect.SliceOf(ftype.Out(0)),
            val.Len(), val.Len())

    var part = float32(val.Len()/factor)
    
    c := make(chan int)
    for i := 0; i < factor; i++ {
        go func(i int) {
            for k := int(part)*i; k < int(part)*(i+1); k++ {
                out.Index(k).Set(reflect.ValueOf(fn(val.Index(k).Interface())))
            }
            c <- i
        }(i)
    }
    
    for i := 0; i < factor; i++ {
        k := <- c
        fmt.Printf("%d part finished ...\n", k)
    }

    return out
}
