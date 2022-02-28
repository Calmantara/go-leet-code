package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	for i := 0; i < 1000; i++ {
		if root.Val == val {
			return root
		}
		if val >= root.Val && root.Right != nil {
			root = root.Right
		} else if val < root.Val && root.Left != nil {
			root = root.Left
		} else {
			break
		}
	}
	return nil
}
