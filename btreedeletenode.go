package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		root = BTreeTransplant(root, node, node.Left)
	} else {
		succ := BTreeMin(node.Right)
		if succ.Parent != node {
			root = BTreeTransplant(root, succ, succ.Right)
			succ.Right = node.Right
			succ.Right.Parent = succ
		}
		root = BTreeTransplant(root, node, succ)
		succ.Left = node.Left
		succ.Left.Parent = succ
	}
	return root
}
