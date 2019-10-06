/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
    if findVal(root2, (target - root1.Val)) {
        return true
    }
    flag := 0
    if root1.Left != nil {
        if twoSumBSTs(root1.Left, root2, target){
            flag++
        }
    }
    if root1.Right != nil {
        if twoSumBSTs(root1.Right, root2, target){
            flag++
        }
    }
    if flag > 0 {
        return true
    } else {
        return false
    }
}

func findVal(root *TreeNode, target int) bool {
    if root.Val == target {
        return true
    }
    flag := 0
    if root.Left != nil {
        if findVal(root.Left, target) {
            flag++
        }
    }
    if root.Right != nil {
        if findVal(root.Right, target){
            flag++
        }
    }
    if flag > 0 {
        return true
    } else {
        return false
    }
}
