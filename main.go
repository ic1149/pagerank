package main

import "fmt"

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
		for k := range prs {
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
			prs[k] = 1 - d + sum
			fmt.Printf("%v:%v", k, prs[k])
		}

	}

}
