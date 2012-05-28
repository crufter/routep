routep
======

The Routep Go package is an overly simplistic way of extracting data from URLs.

```
package main

import(
	"github.com/opesun/routep"
	"fmt"
)

func main() {
	m, ok := routep.Comp("/categories/{category}", "/categories/animals")
	fmt.Println(m, ok)
}
```

The above example will output:
```
map["category": "animals"] true
```
