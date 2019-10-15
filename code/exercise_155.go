type MinStack struct {
    values []int
    mins []int
    index int
    minidx int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    ms := new(MinStack)
    return *ms
}


func (this *MinStack) Push(x int)  {
    //fmt.Printf("%v ------> Push : %d\n", *this, x)
    if len(this.values) <= this.index {
        this.values = append(this.values, x)
    } else {
        this.values[this.index] = x
    }
    if this.minidx <= 0 || this.values[this.mins[this.minidx-1]] >= x {
        if len(this.mins) <= this.minidx {
            this.mins = append(this.mins, this.index)
        } else {
            this.mins[this.minidx] = this.index
        }
        this.minidx++
    }
    this.index++
}


func (this *MinStack) Pop()  {
    //fmt.Printf("%v ------> Pop\n", *this)
    if this.index>0 {
        
        if this.minidx>0 && this.values[this.index-1] == this.values[this.mins[this.minidx-1]] {
            this.minidx--
        }
        this.index--
    }
}


func (this *MinStack) Top() int {
    return this.values[this.index-1]
}


func (this *MinStack) GetMin() int {
    return this.values[this.mins[this.minidx-1]]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */