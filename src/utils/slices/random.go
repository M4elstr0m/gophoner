package gp_slices

import "math/rand"

func RandomFrom[S ~[]E, E comparable](s S) E {
	return s[rand.Intn(len(s))]
}
