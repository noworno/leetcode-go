/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
    res := L + R
    bst(root, &res, L, R)
    return res
}

func bst(root *TreeNode, r *int, L,R int) {
    if root.Val > L && root.Left != nil {
        bst(root.Left, r, L, R)
    }
    if root.Val > L && root.Val < R {
         *r = *r + root.Val
    }
    if root.Val < R && root.Right != nil {
        bst(root.Right, r, L, R)
    }
}