package _slices

func Inter[S ~[]E, E comparable](a S, b S) S {
	set := make(map[E]bool, len(a))
	for _, v := range a {
		set[v] = true
	}
	out := make(S, 0, len(b))
	for _, v := range b {
		if set[v] {
			out = append(out, v)
		}
	}
	return out
}
