package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}
	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}
	it := l
	for it.Next != nil && it.Next.Data < data_ref {
		it = it.Next
	}
	newNode.Next = it.Next
	it.Next = newNode
	return l
}
