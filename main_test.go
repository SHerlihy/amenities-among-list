package main

import (
	"testing"
)

type TestInput struct {
    title string
	desiredAmenities []string
	amenitiesByBlock [][]string
    expectedIdx int
}

func TestDistances(t *testing.T) {
	desiredAmenities := []string{"chippy", "shop", "park"}

	amenitiesByBlock := [][]string{[]string{"shop", "chippy"}, []string{"chippy"}, []string{"chippy", "park"}}

	testInputs := []TestInput{
		TestInput{"min example", desiredAmenities[0:3], amenitiesByBlock, 1},
	}

	for _, testInput := range testInputs {
		t.Run(testInput.title, func(t *testing.T) {
			shortest := Main(testInput.desiredAmenities, testInput.amenitiesByBlock)
			if shortest != testInput.expectedIdx {
				t.Fatalf(`Expected %v but got %v`, testInput.expectedIdx, shortest)
			}
		})
	}
}
