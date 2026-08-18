package _slices

func Merge[S ~[]E, E any](slices ...S) S {
	var total int
	for _, s := range slices {
		total += len(s)
	}

	out := make(S, 0, total)
	for _, s := range slices {
		out = append(out, s...)
	}

	return out
}
