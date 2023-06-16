package main

import (
	"math"
)

type PotentialEl struct {
	index     int
	amenities []string
}

func Main(desiredAmenities []string, amenitiesByBlock [][]string) int {
	var potentialEls []PotentialEl
	var completeIdx []int

	for poolIdx, poolEls := range amenitiesByBlock {

		includedDesired := includedElements(poolEls, desiredAmenities)

		if len(includedDesired) == 0 {
			continue
		} else {
			potentialEle := PotentialEl{poolIdx, poolEls}
			potentialEls = append(potentialEls, potentialEle)
		}
	}

	for _, potentialEl := range potentialEls {
		elIdx := potentialEl.index

		endingIdx := -1

		remainingDesired := desiredAmenities

		for _, matchVal := range potentialEls {
			matchIdx := matchVal.index
			matchAmenities := matchVal.amenities

			updatedDesired := removeIncluded(matchAmenities, remainingDesired)

			remainingDesired = updatedDesired

			if len(remainingDesired) == 0 {
				endingIdx = matchIdx
				break
			}
		}

		if endingIdx != -1 {
			midOffset := math.Floor(float64((endingIdx - elIdx) / 2))
			midIdx := elIdx + int(midOffset)

			completeIdx = append(completeIdx, midIdx)
		}

	}

	return getLowest(completeIdx)
}

func getLowest(s1 []int) int {
	var m int
	for i, e := range s1 {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

func includedElements(subjects []string, targets []string) []string {
	var included []string

	for _, targetEl := range targets {
		containsTarget := contains(subjects, targetEl)

		if containsTarget != -1 {
			included = append(included, targetEl)
		}
	}

	return included
}

func removeIncluded(subjects []string, targets []string) []string {
	remains := targets

	for _, subjectEl := range subjects {
		containsTarget := contains(remains, subjectEl)
		if containsTarget != -1 {
			remains = remove(remains, containsTarget)
		}
	}

	return remains
}

func contains(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
