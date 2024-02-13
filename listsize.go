package piscine

func ListPushFront(l *List, data interface{}) {
	node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = node
		return
	}
	node.Next = l.Head
	l.Head = node
}

func ListSize(l *List) int {
	size := 0
	current := l.Head
	if current != nil {
		size++
		current = current.Next
	}
	return size
}
