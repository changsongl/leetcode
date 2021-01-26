package main

// 非命中数字，拷贝到前面
func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}

		if i != j {
			nums[j] = nums[i]
		}
		j++
	}
	return j
}

// 命中的数字替换到后面
func removeElementEnd(nums []int, val int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		if nums[start] != val {
			start++
			continue
		}

		for start <= end {
			if nums[end] != val {
				nums[start] = nums[end]
				nums[end] = val
				end--
				break
			}
			end--
		}
	}
	return start
}
