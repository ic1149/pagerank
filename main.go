package main

import (
	"fmt"
	"math"
	"sort"
)

func roundFloat32(f float32, precision int) float32 {
	shift := math.Pow10(precision)
	return float32(math.Round(float64(f)*shift) / shift)
}

func main() {
	matrix := map[string]map[string]bool{
		"a": {"a": false, "b": true, "c": true, "d": false},
		"b": {"a": false, "b": false, "c": true, "d": false},
		"c": {"a": true, "b": false, "c": false, "d": false},
		"d": {"a": false, "b": false, "c": true, "d": false},
	}
	prs := map[string]float32{
		"a": 1.0,
		"b": 1.0,
		"c": 1.0,
		"d": 1.0,
	}

	const d float32 = 0.85

	for range 10 {
		keys := make([]string, 0, len(prs))
		for key := range prs {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, k := range keys {
			var sum float32 = 0

			for i, w := range matrix {
				if w[k] == true {
					var c float32 = 0
					for _, x := range w {
						if x == true {
							c += 1
						}
					}
					sum += d * prs[i] / c
				}
			}
			prs[k] = roundFloat32(1-d+sum, 3)
			fmt.Printf("%v:%v ", k, prs[k])
		}
		fmt.Print("\n")

	}

}
