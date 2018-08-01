package stringutil_test

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

func ExampleReverse() {
	fmt.Println(stringutil.Reverse("abcdefghijklmnopqrstuvwxyz"))
	// Output: olleh
}
