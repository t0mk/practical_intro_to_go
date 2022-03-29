package main

import (
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	// Test for average of a float slice
	fmt.Println("Test for average of a float slice")
	example := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if average := floatSliceAverage(example); average != 5.5 {
		t.Errorf("Average of %v is %f, expected 5.5", example, average)
	}
}
