package main

import (
	"fmt"
	"strconv"
)

// V1 convert [r,g,b] tuple into [r,g,b,hex] tuple
func V1(r, g, b, i int) ([]string, error) {
	return []string{
		strconv.Itoa(r),
		strconv.Itoa(g),
		strconv.Itoa(b),
		fmt.Sprintf("#%02X%02X%02X", r, g, b),
	}, nil
}
