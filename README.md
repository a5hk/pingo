Simple wrapper around Linux `ping` command.

# Usage
```GO
package main

import (
	"fmt"
	"github.com/ashksh/pingo"
)

func main() {
	stats, _ := pingo.Ping("google.com", "-c", "1")
	fmt.Printf("%+v\n", stats)
}
```

`Output: {Transmitted:1 Received:1 Errors:0 Loss:0 Time:0 Min:145.867 Avg:145.867 Max:145.867 Mdev:0}`
