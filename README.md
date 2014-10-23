#parallel

A minimalist go library to simplify concurrency programming in Go. It is non-trival to write 
concurrency program. This library aims to simplify the process of writing 
parallel/concurrency program.

## Installation

```go
go get github.com/codingneo/parallel
```

## Quick Examples

```go
import (
	"github.com/codingneo/parallel"
)

func main() {
    data := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

    parallel.Iterator{0,15,1}.For(
        func(i int) { data[i] = data[i]*2 } )
}
```

## Documentation
### Struct
* [`Iterator`](#Iterator)

### APIs
* [`For`](#For)

<a name="Iterator" />
### Iterator
```go
type Iterator struct {
	Start int
	End int
	Step int
}
```

Iterator struct defines the necessary information about a for loop. 

__Arguments__

* `Start` - It defines the starting index of a for loop.
* `End` - It defines the ending index (excluded) of a for loop.
* `Step` - It defines the step size of a for loop (Default is 1).


<a name="For" />
### For
```go
func (iter Iterator) For(block func(int))
```

For function is a method of an Iterator struct. To call the For function, you 
need to create an Iterator struct (very often as a struct literal) and invoke 
For method using a chain invokation. The single parameter is a function with 
single input parameter of type int.

__Arguments__

* `block` - The actually function to execute with index i. It is a function 
closure which can manipulate any object in the context.
