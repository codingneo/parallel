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
```

## Documentation
### Struct
* [`Iterator`](#Iterator)

### APIs
* [`For`](#For)
* [`Map`](#Map)
* [`FlatMap`](#FlatMap)
* [`Filter`](#Filter)
* [`Reduce`](#Reduce)

<a name="Iterator" />
### Iterator
```go
type Iterator struct {
	Start int
	End int
	Step int
}
```

Produces a new array of values by mapping each value in `arr` through
the `iterator` function. The `iterator` is called with an item from `arr` and a
callback for when it has finished processing. Each of these callback takes 2 arguments: 
an `error`, and the transformed item from `arr`. If `iterator` passes an error to his 
callback, the main `callback` (for the `map` function) is immediately called with the error.

Note, that since this function applies the `iterator` to each item in parallel,
there is no guarantee that the `iterator` functions will complete in order. 
However, the results array will be in the same order as the original `arr`.

__Arguments__

* `arr` - An array to iterate over.
* `iterator(item, callback)` - A function to apply to each item in `arr`.
  The iterator is passed a `callback(err, transformed)` which must be called once 
  it has completed with an error (which can be `null`) and a transformed item.
* `callback(err, results)` - A callback which is called when all `iterator`
  functions have finished, or an error occurs. Results is an array of the
  transformed items from the `arr`.

__Example__

```js
async.map(['file1','file2','file3'], fs.stat, function(err, results){
    // results is now an array of stats for each file
});
```
