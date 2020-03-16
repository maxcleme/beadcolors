package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// GenFunc represent a translation between raw tuple (r,g,b) into something else
type GenFunc func(r, g, b int) ([]string, error)

var all = map[string]GenFunc{
	"/v1": V1,
}

func main() {
	if len(os.Args) < 3 {
		panic(fmt.Errorf("main: missing params"))
	}

	raw := os.Args[1]
	output := os.Args[2]
	err := filepath.Walk(raw, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		r := csv.NewReader(f)
		lines, err := r.ReadAll()
		if err != nil {
			return err
		}
		for v, gen := range all {
			_ = os.Mkdir(filepath.Join(output, v), os.ModePerm)
			dest, err := os.Create(filepath.Join(output, v, info.Name()))
			if err != nil {
				return err
			}

			w := csv.NewWriter(dest)
			for _, l := range lines {
				// Raw .csv contains only 5 columns [ref, name, r, g, b]
				if len(l) != 5 {
					return fmt.Errorf("main: invalid raw format")
				}
				ref := l[0]
				name := l[1]
				r, g, b, err := parseRGB(l[2], l[3], l[4])
				if err != nil {
					return err
				}

				ss, err := gen(r, g, b)
				if err != nil {
					return err
				}

				if err := w.Write(append([]string{ref, name}, ss...)); err != nil {
					return err
				}
			}
			w.Flush()
			fmt.Printf("Generation\t%s\t%s\tOK\n", v, info.Name())

		}
		return nil
	})
	if err != nil {
		panic(fmt.Errorf("main: cannot walk through raw directory: %w", err))
	}
}

func parseRGB(r, g, b string) (int, int, int, error) {
	rc, err := strconv.Atoi(r)
	if err != nil {
		return 0, 0, 0, err
	}
	gc, err := strconv.Atoi(g)
	if err != nil {
		return 0, 0, 0, err
	}
	bc, err := strconv.Atoi(b)
	if err != nil {
		return 0, 0, 0, err
	}
	if rc < 0 || rc > 255 {
		return 0, 0, 0, fmt.Errorf("invalid red component : %d", rc)
	}
	if gc < 0 || gc > 255 {
		return 0, 0, 0, fmt.Errorf("invalid green component : %d", rc)
	}
	if bc < 0 || bc > 255 {
		return 0, 0, 0, fmt.Errorf("invalid blue component : %d", rc)
	}

	return rc, gc, bc, nil
}
