package main

func isMonotonic(A []int) bool {
	less, large := true, true

	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			large = false
		} else if A[i] < A[i+1] {
			less = false
		}
	}

	return less || large
}
