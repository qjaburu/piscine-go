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
