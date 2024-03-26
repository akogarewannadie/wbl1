package main

import (
	"fmt"
	"math"
)

func main() {

	groups := make(map[int][]float64)

	step := 10

	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, temp := range temperatures {
		groupKey := int(math.Floor(temp/float64(step))) * step
		groups[groupKey] = append(groups[groupKey], temp)
	}

	for key, values := range groups {
		fmt.Printf("Group %d°C: %v\n", key, values)
	}
}
