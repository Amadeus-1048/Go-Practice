package error_learn

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {
	res, err := Sqrt(-1)
	fmt.Println(res)
	fmt.Println(err)
}
