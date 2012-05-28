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
	m, err := routep.Comp("/categories/{category}", "/categories/animals")
	fmt.Println(m, err)
}
```

The above example will output (not counting the lack of '"'):
```
map["category": "animals"] ""
```
