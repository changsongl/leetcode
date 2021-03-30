package main

//import "fmt"
//
//func main() {
//	f := Constructor()
//	f.AddNum(1)
//	f.Print()
//	f.AddNum(3)
//	f.Print()
//	f.AddNum(5)
//	fmt.Print(f.FindMedian())
//	f.Print()
//	f.AddNum(2)
//	fmt.Print(f.FindMedian())
//	f.Print()
//}
//
//type MedianFinder struct {
//	nums     []int
//	midIndex int
//	isEven   bool
//}
//
//func (this *MedianFinder) Len() int {
//	return len(this.nums)
//}
//
///** initialize your data structure here. */
//func Constructor() MedianFinder {
//	return MedianFinder{}
//}
//
//func (this *MedianFinder) AddNum(num int) {
//	if this.Len() == 0 {
//		this.midIndex = 0
//		this.nums = append(this.nums, num)
//		return
//	}
//
//	if this.isEven {
//		this.midIndex++
//	}
//	this.isEven = !this.isEven
//
//	// 二分插入
//	index := this.findInsertIndex(num)
//
//	var nums []int
//	nums = append(nums, this.nums[0:index]...)
//	nums = append(nums, num)
//	nums = append(nums, this.nums[index:]...)
//	this.nums = nums
//}
//
//func (this *MedianFinder) findInsertIndex(num int) int {
//	l, r := 0, this.Len()-1
//	for l <= r {
//		m := (l + r) / 2
//		if this.nums[m] == num {
//			for m < this.Len() && this.nums[m] == num {
//				m++
//			}
//			return m - 1
//		} else if this.nums[m] > num {
//			r = m - 1
//		} else {
//			l = m + 1
//		}
//	}
//	return r + 1
//}
//
//func (this *MedianFinder) Print() {
//	for _, num := range this.nums {
//		fmt.Printf("%d ", num)
//	}
//	fmt.Printf("\nindex: %d, isEven: %t\n", this.midIndex, this.isEven)
//}
//
//func (this *MedianFinder) FindMedian() float64 {
//	if this.isEven {
//		return (float64(this.nums[this.midIndex]) + float64(this.nums[this.midIndex+1])) / 2
//	} else {
//		return float64(this.nums[this.midIndex])
//	}
//}
