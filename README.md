#parallel

A go library to simplify concurrency programming. It is non-trival to write 
concurrency program. This library aims to simplify the process of writing 
parallel/concurrency program.

## Quick Examples

```go
import (
	"fmt"
	"www.github.com/codingneo/parallel"
)

func main() {
	fmt.Println("Hello World")

	data := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  parallel.For(parallel.Iterator{0, 20, 2},
                func(i int) { fmt.Println(data[i]) })
}

