package operations

import (
	"fmt"
	"strconv"
)

func Sum(a, b string) string {
	an, _ := strconv.ParseInt(a, 10, 32)
	bn, _ := strconv.ParseInt(b, 10, 32)
	return fmt.Sprintf("%d", an+bn)
}
