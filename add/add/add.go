package add

import "log"

// Add func returns the sum of two numbers
func Add(a, b int) int {
	log.Println("Adding ", a, " and ", b)
	return a + b
}
