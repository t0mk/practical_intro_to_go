package main

import (
	"fmt"

	"github.com/t0mk/gokativ"
)

// show
//  - anatomy of a .go file
//  - declared but not used
//  - slice (float slice)
//  - function
//  - map
//  - range operator for map
//  - Pause and ask

func floatSliceAverage(slice []float64) float64 {
	var sum float64
	for _, value := range slice {
		sum += value
	}
	return sum / float64(len(slice))
}

// m = {"a": 1, "b": 2, "c": 3}

func main() {
	fmt.Println(gokativ.Vokativ("Tomáš"))
	fmt.Println("Hello World")
	example := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(floatSliceAverage(example))
	exampleMap := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range exampleMap {
		fmt.Println(key, value)
	}
	m := map[string]map[string]int{
		"a": {"a": 1, "b": 2, "c": 3},
		"b": {"a": 1, "b": 2, "c": 3},
	}
	fmt.Println(m)
}
