package piscine

func Swap(a *int, b *int) {
	char := *a

	*a = *b

	*b = char
}
