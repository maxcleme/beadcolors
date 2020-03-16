package main

import (
	"fmt"
	"strconv"
)

// V1 convert [r,g,b] tuple into [r,g,b,hex] tuple
func V1(r, g, b int) ([]string, error) {
	return []string{
		strconv.Itoa(r),
		strconv.Itoa(g),
		strconv.Itoa(b),
		fmt.Sprintf("#%X%X%X", r, g, b),
	}, nil
}
