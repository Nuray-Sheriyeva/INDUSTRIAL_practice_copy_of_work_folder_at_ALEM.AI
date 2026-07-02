package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if found := BTreeSearchItem(root.Left, elem); found != nil {
		return found
	}
	return BTreeSearchItem(root.Right, elem)
}
