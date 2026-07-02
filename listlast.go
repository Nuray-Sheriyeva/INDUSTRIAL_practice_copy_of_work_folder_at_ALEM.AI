package piscine

func ListLast(l *List) interface{} {
	n := l.Tail
	if n == nil {
		return nil
	} else {
		dt := n.Data
		return dt
	}
}
