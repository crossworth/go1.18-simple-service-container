### Simple example of a service container using GO 1.18 generics

```go
package main

import (
	"fmt"
)

type Database struct {
	val string
}

func main() {
	db := &Database{
		val: "test string",
	}

	// store the dependency
	Put(db)

	// get the dependency
	dbDep, found := Get[*Database]()
	if !found {
		fmt.Printf("could not find the dependency\n")
	} else {
		fmt.Printf("%#v\n", dbDep)
	}
}
```