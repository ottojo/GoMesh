package main

import "math/rand"

func getRandomWithout(lowerBorderInclusive, upperBorderExclusive, not int) int {
	var random int = rand.Intn(upperBorderExclusive-lowerBorderInclusive) + lowerBorderInclusive
	for random == not {
		random = rand.Intn(upperBorderExclusive-lowerBorderInclusive) + lowerBorderInclusive
	}
	return random
}
