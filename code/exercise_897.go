/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    return  dfs(root)
}

func dfs(root *TreeNode) Node {
    head := new(TreeNode)
    s := new(Stack)
    s.Push(root)
    for s.Size()>0 {
        v := s.Top()
        if v.Right!=nil {
            s.Push(v.Right)
            v.Right = nil
            continue
        }
        s.Pop()
        v.Right = head.Right
        head.Right = v
        if v.Left!=nil{
            s.Push(v.Left)
            v.Left = nil
        }
    }
    return head.Right
}

type Node *TreeNode
type Stack struct {
	arr   []Node
	index int
}

func (s *Stack) Pop() Node{
	if s.index > 0 {
		s.index--
		return s.arr[s.index]
	}
	return nil
}

func (s *Stack) Push(x Node) {
	if len(s.arr) <= s.index {
		s.arr = append(s.arr, x)
	} else {
		s.arr[s.index] = x
	}
	s.index++
}

func (s *Stack) Top() Node {
    if s.index > 0 {
		return s.arr[s.index-1]
	}
	return nil
}

func (s *Stack) Size() int {
	return s.index
}