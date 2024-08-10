package generator

import "math/rand"

func NewPageReferences(pagesNumber, referencesLength int) []int {
	pageReferences := make([]int, referencesLength)

	for range referencesLength {
		pageReferences = append(pageReferences, rand.Int()%pagesNumber)
	}

	return pageReferences
}
