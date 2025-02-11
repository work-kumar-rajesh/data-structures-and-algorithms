package main

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

// Implement the MinStack class:

// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.

// solution
//Main idea is that store global minimum with each value


type pair struct {
    value int
    currMin int
}

type MinStack struct {
    stack []pair
}


func Constructor() MinStack {
    stack := make([]pair,0)
    return MinStack{stack}
}


func (this *MinStack) Push(val int)  {
    if len(this.stack) == 0 {
        this.stack = append(this.stack,pair{val,val})
    }else{
        top := this.stack[len(this.stack)-1]
        this.stack = append(this.stack,pair{val,min(val,top.currMin)})
    }
}


func (this *MinStack) Pop()  {
    if len(this.stack) > 0 {
        this.stack = this.stack[:len(this.stack)-1]
    }
}