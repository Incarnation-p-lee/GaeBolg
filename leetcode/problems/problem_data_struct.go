package problems

type Interval struct {
    Start int
    End   int
}

type ListNode struct {
    Val int
    Next *ListNode
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type minHeap struct {
	Size int
	Array []*ListNode
}

