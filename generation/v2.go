package main

import (
	"fmt"
	"math"
)

// V2 convert [r,g,b] tuple into [rgb_r,rgb_g,rgb_b,hsl_h,hsl_s,hsl_l,lab_l,lab_a,lab_b] tuple
func V2(rgbR, rgbG, rgbB int) ([]string, error) {
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

func xyz(r, g, b int) (float64, float64, float64, error) {
	rn := float64(r) / 255
	gn := float64(g) / 255
	bn := float64(b) / 255

	if rn > 0.04045 {
		rn = math.Pow((rn+0.055)/1.055, 2.4)
	} else {
		rn = rn / 12.92
	}

	if gn > 0.04045 {
		gn = math.Pow((gn+0.055)/1.055, 2.4)
	} else {
		gn = gn / 12.92
	}

	if bn > 0.04045 {
		bn = math.Pow((bn+0.055)/1.055, 2.4)
	} else {
		bn = bn / 12.92
	}

	rn = rn * 100
	gn = gn * 100
	bn = bn * 100

	xyzX := rn*0.4124 + gn*0.3576 + bn*0.1805
	xyzY := rn*0.2126 + gn*0.7152 + bn*0.0722
	xyzZ := rn*0.0193 + gn*0.1192 + bn*0.9505
	return xyzX, xyzY, xyzZ, nil
}

func lab(r, g, b int) (float64, float64, float64, error) {
	x, y, z, err := xyz(r, g, b)
	if err != nil {
		return 0, 0, 0, err
	}

	x = x / 95.047
	y = y / 100.0
	z = z / 108.883

	if x > 0.008856 {
		x = math.Pow(x, 1.0/3.0)
	} else {
		x = (7.787 * x) + (16 / 116)
	}

	if y > 0.008856 {
		y = math.Pow(y, 1.0/3.0)
	} else {
		y = (7.787 * y) + (16 / 116)
	}

	if z > 0.008856 {
		z = math.Pow(z, 1.0/3.0)
	} else {
		z = (7.787 * z) + (16 / 116)
	}

	labL := (116 * y) - 16
	labA := 500 * (x - y)
	labB := 200 * (y - z)

	return labL, labA, labB, nil
}

func hsl(r, g, b int) (float64, float64, float64, error) {
	rn := float64(r) / 255
	gn := float64(g) / 255
	bn := float64(b) / 255

	min := math.Min(rn, math.Min(gn, bn))
	max := math.Max(rn, math.Max(gn, bn))

	hslL := (min + max) / 2

	if min == max {
		return 0, 0, hslL, nil
	}

	hslS := 0.0
	if hslL > 0.5 {
		hslS = (max - min) / (2 - max - min)
	} else {
		hslS = (max - min) / (max + min)
	}

	hslH := 0.0
	switch max {
	case rn:
		hslH = (gn - bn) / (max - min)
	case gn:
		hslH = 2 + (bn-rn)/(max-min)
	case bn:
		hslH = 4 + (rn-gn)/(max-min)
	}

	hslH *= 60
	if hslH < 0 {
		hslH += 360
	}
	return hslH, hslS, hslL, nil
}
