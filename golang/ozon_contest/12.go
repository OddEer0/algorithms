package ozoncontest

import (
	"fmt"
	"math"
)

func Task12Handle() float64 {
	var count, percent int
	fmt.Scan(&count, &percent)
	remind := 0.0
	for i := 0; i < count; i++ {
		var tmp int
		fmt.Scan(&tmp)
		pr := float64(tmp*percent) / 100.0
		pr -= math.Floor(pr)
		remind += pr
	}
	return remind
}

func Task12() {
	var count int
	fmt.Scan(&count)
	results := make([]float64, count)

	for i := 0; i < count; i++ {
		results[i] = Task12Handle()
	}

	for i := 0; i < count; i++ {
		fmt.Printf("%.2f", results[i])
		if i != count-1 {
			fmt.Print("\n")
		}
	}
}
