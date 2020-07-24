package main

import (
	"fmt"
	"strconv"
)

// V1 convert [r,g,b] tuple into [r,g,b,hex] tuple
// Index is the identifier in each colour pallette item.
func V1(r, g, b, index int) ([]string, error) {
	return []string{
		strconv.Itoa(r),
		strconv.Itoa(g),
		strconv.Itoa(b),
		fmt.Sprintf("#%02X%02X%02X", r, g, b),
	}, nil
}
