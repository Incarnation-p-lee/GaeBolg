package problems

import "fmt"

type MinStack struct {
	Stack []int
	MinIndex []int    
}

func minConstructor() MinStack {
	mstack := new(MinStack)

	return *mstack
}

func (this *MinStack) Push(x int)  {
	if this == nil {
		return
	} 

	if len(this.Stack) == 0 {
		this.Stack = append(this.Stack, x)
		this.MinIndex = append(this.MinIndex, 0)
	} else {
		i := this.MinIndex[len(this.MinIndex) - 1]
		this.Stack = append(this.Stack, x)

		if this.Stack[i] >= x {
			this.MinIndex = append(this.MinIndex, len(this.Stack) - 1)
		}
	}
}

func (this *MinStack) Pop() {
	if this == nil {
		return
	} else if len(this.Stack) == 0 {
		return
	}

	i := len(this.Stack) - 1
	this.Stack = this.Stack[:len(this.Stack) - 1]

	if this.MinIndex[len(this.MinIndex) - 1] == i {
		this.MinIndex = this.MinIndex[:len(this.MinIndex) - 1]
	}
}

func (this *MinStack) Top() int {
	if this == nil {
		return -1
	} else if len(this.Stack) == 0 {
		return -1
	}

	return this.Stack[len(this.Stack) - 1]
}

func (this *MinStack) GetMin() int {
	if this == nil {
		return -1
	} else if len(this.Stack) == 0 {
		return -1
	}

	return this.Stack[this.MinIndex[len(this.MinIndex) - 1]]
}

func MinStackSample() {
	minStack := minConstructor()

	fmt.Printf("<115> ")

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	fmt.Printf("%d ", minStack.GetMin())

	minStack.Pop()

	fmt.Printf("%d ", minStack.Top())
	fmt.Printf("%d ", minStack.GetMin())
	fmt.Println()
}

