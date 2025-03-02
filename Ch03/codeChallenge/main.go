// Write your answer here, and then test your code.

package main

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// slicesToObjects() returns a slice of Color objects

func slicesToObjects(colorNames []string, hexValues []int) []Color {
	// Your code goes here.
	// Create a slice of Color object
	var colors []Color = make([]Color, 0)

	for i := range colorNames {
		colors = append(colors, Color{colorNames[i], hexValues[i]})
	}

	return colors
}

type Color struct {
	Name string
	Hex  int
}
