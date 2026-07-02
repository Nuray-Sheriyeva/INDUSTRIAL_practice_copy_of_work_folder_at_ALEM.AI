package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	this := l
	for i := 0; i < pos; i++ {
		if this.Next != nil {
			this = this.Next
		} else {
			return nil
		}
	}
	return this
}
