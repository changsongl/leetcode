package main

// 速度91%, 内存100%
func main() {
	n := Constructor([]*NestedInteger{&NestedInteger{}})
	n.HasNext()
}

type NestedIterator struct {
	sub []SubIterator
}

type SubIterator struct {
	ni []*NestedInteger
	i  int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		sub: []SubIterator{
			{nestedList, 0},
		},
	}
}

func (this *NestedIterator) Next() int {
	if this.HasNext() {
		integers := this.sub[len(this.sub)-1]
		num := integers.ni[integers.i].GetInteger()
		integers.i++
		this.sub[len(this.sub)-1] = integers
		return num
	}

	return 0
}

func (this *NestedIterator) HasNext() bool {
	for i := len(this.sub) - 1; i >= 0; i-- {
		integerSlice := this.sub[i]
		if len(integerSlice.ni) <= integerSlice.i {
			this.sub = this.sub[:i]
			continue
		}

		integer := integerSlice.ni[integerSlice.i]
		if integer.IsInteger() {
			return true
		} else {
			integerSlice.i++
			this.sub[i] = integerSlice
			this.sub = append(this.sub, SubIterator{integer.GetList(), 0})
			i = len(this.sub)
		}
	}

	return false
}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { return false }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { return 0 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { return []*NestedInteger{} }
