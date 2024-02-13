package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

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
	for current != nil {
		size++
		current = current.Next
	}
	return size
}
