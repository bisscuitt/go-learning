// Package dog implements math functions for dogs
package dog

// Years converts human years to dog years and returns an int
func Years(y int) int {
	if y < 0 {
		return 0
	}
	return y * 7
}
