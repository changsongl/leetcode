package main

type MedianFinder struct {
	min *MinHeap
	max *MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		min: NewMinHeap(),
		max: NewMaxHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.min.Size() != this.max.Size() {
		this.min.PushHeap(num)
		this.max.PushHeap(this.min.PopHeap())
	} else {
		this.max.PushHeap(num)
		this.min.PushHeap(this.max.PopHeap())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.min.Size() != this.max.Size() {
		return float64(this.min.TopHeap())
	} else {
		return (float64(this.max.TopHeap()) + float64(this.min.TopHeap())) / 2
	}
}

// MaxHeap 最大堆结构体
type MaxHeap struct {
	data []int
	size int
}

// NewMaxHeap 建立一个最大堆
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data: make([]int, 16)[:0],
		size: 0,
	}
}

// PushHeap 向最大堆压入一个数字
func (mh *MaxHeap) PushHeap(val int) {
	if mh.size == 0 {
		mh.data = append(mh.data, val)
		mh.size++
		return
	}
	mh.data = append(mh.data, val)
	mh.size++
	mh.shiftUp()
}

// PopHeap 用来弹出堆顶元素
func (mh *MaxHeap) PopHeap() int {
	if mh.size == 0 {
		return -1
	}
	res := mh.data[0]
	mh.size--
	mh.data[0] = mh.data[mh.size]
	mh.data = mh.data[:mh.size]
	if mh.size <= 1 {
		return res
	}
	mh.shiftDown()
	return res
}

// 把最后一个数字依次向父节点挪动
// 维持最大堆的特性
func (mh *MaxHeap) shiftUp() {
	idx := len(mh.data) - 1
	for idx >= 0 {
		fidx := (idx - 1) / 2
		if mh.data[fidx] < mh.data[idx] {
			mh.data[fidx], mh.data[idx] = mh.data[idx], mh.data[fidx]
			idx = fidx
		} else {
			break
		}
	}
}

// 把第一个数字向下移动
// 维持最大堆
func (mh *MaxHeap) shiftDown() {
	idx := 0
	for idx < len(mh.data) {
		lidx := idx*2 + 1
		ridx := idx*2 + 2
		largeidx := idx
		if lidx < mh.size && mh.data[lidx] > mh.data[largeidx] {
			largeidx = lidx
		}
		if ridx < mh.size && mh.data[ridx] > mh.data[largeidx] {
			largeidx = ridx
		}
		if largeidx == idx {
			break
		}
		mh.data[largeidx], mh.data[idx] = mh.data[idx], mh.data[largeidx]
		idx = largeidx
	}
}

// TopHeap 返回堆顶元素但不弹出
func (mh *MaxHeap) TopHeap() int {
	if mh.size == 0 {
		return -1
	}
	return mh.data[0]
}

// Size 返回堆的大小
func (mh *MaxHeap) Size() int {
	return mh.size
}

// MinHeap 最小堆结构体
type MinHeap struct {
	data []int
	size int
}

// NewMinHeap 建立一个最小堆
func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: make([]int, 16)[:0],
		size: 0,
	}
}

// PushHeap 向最小堆压入一个数字
func (mh *MinHeap) PushHeap(val int) {
	if mh.size == 0 {
		mh.data = append(mh.data, val)
		mh.size++
		return
	}
	mh.data = append(mh.data, val)
	mh.size++
	mh.shiftUp()
}

// PopHeap 用来弹出堆顶元素
func (mh *MinHeap) PopHeap() int {
	if mh.size == 0 {
		return -1
	}
	res := mh.data[0]
	mh.size--
	mh.data[0] = mh.data[mh.size]
	mh.data = mh.data[:mh.size]
	if mh.size <= 1 {
		return res
	}
	mh.shiftDown()
	return res
}

// 把最后一个数字依次向父节点挪动
// 维持最小堆的特性
func (mh *MinHeap) shiftUp() {
	idx := len(mh.data) - 1
	for idx >= 0 {
		fidx := (idx - 1) / 2
		if mh.data[fidx] > mh.data[idx] {
			mh.data[fidx], mh.data[idx] = mh.data[idx], mh.data[fidx]
			idx = fidx
		} else {
			break
		}
	}
}

// 把第一个数字向下移动
// 维持最小堆
func (mh *MinHeap) shiftDown() {
	idx := 0
	for idx < len(mh.data) {
		lidx := idx*2 + 1
		ridx := idx*2 + 2
		smallidx := idx
		if lidx < mh.size && mh.data[lidx] < mh.data[smallidx] {
			smallidx = lidx
		}
		if ridx < mh.size && mh.data[ridx] < mh.data[smallidx] {
			smallidx = ridx
		}
		if smallidx == idx {
			break
		}
		mh.data[smallidx], mh.data[idx] = mh.data[idx], mh.data[smallidx]
		idx = smallidx
	}
}

// TopHeap 返回堆顶元素但不弹出
func (mh *MinHeap) TopHeap() int {
	if mh.size == 0 {
		return -1
	}
	return mh.data[0]
}

// Size 返回堆的大小
func (mh *MinHeap) Size() int {
	return mh.size
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
