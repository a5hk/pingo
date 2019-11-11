package pingo

import (
	"fmt"
	"github.com/ashksh/pingo"
)

func Example_basic() {
	stats, _ := pingo.Ping("google.com", "-c", "1")
	fmt.Printf("%+v\n", stats)
}
