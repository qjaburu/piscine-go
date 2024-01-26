package piscine

func UltimateDivMod(a *int, b *int) {
	i := *a
	k := *b

	*a = i / k
	*b = i % k
}
