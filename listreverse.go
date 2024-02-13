package piscine

func ListReverse(l *List) {
	var previous, current, Next *NodeL
	current = l.Head
	l.Tail = l.Head
	for current != nil {
		Next = current.Next
		current.Next = previous
		previous = current
		current = Next
	}
	l.Head = previous
}
