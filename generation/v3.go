package main

import (
	"fmt"
)

// V3 convert [r,g,b] tuple into [rgb_r,rgb_g,rgb_b,hsl_h,hsl_s,hsl_l,lab_l,lab_a,lab_b] tuple
func V3(rgbR, rgbG, rgbB int) ([]string, error) {
	hslH, hslS, hslL, err := hsl(rgbR, rgbG, rgbB)
	if err != nil {
		return nil, err
	}
	labL, labA, labB, err := lab(rgbR, rgbG, rgbB)
	if err != nil {
		return nil, err
	}
	return []string{
		fmt.Sprintf("%d", rgbR),
		fmt.Sprintf("%d", rgbG),
		fmt.Sprintf("%d", rgbB),
		fmt.Sprintf("%.2f", hslH),
		fmt.Sprintf("%.2f", hslS),
		fmt.Sprintf("%.2f", hslL),
		fmt.Sprintf("%.2f", labL),
		fmt.Sprintf("%.2f", labA),
		fmt.Sprintf("%.2f", labB),
	}, nil
}
