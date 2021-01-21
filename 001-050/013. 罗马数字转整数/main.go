package main

// 执行结果：通过显示详情执行用时：0 ms, 在所有 Go 提交中击败了 `100.00%` 的用户
// 内存消耗：3.1 MB, 在所有 Go 提交中击败了 `46.02%` 的用户
// 一次通过，空间换时间。
func romanToInt(s string) int {
	numMap := map[int32]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	prev := map[int32]int32{
		'V': 'I',
		'X': 'I',
		'L': 'X',
		'C': 'X',
		'D': 'C',
		'M': 'C',
	}
	prevCh := '0'
	num := 0

	for _, b := range s {
		if prev[b] == prevCh {
			num += numMap[b] - 2*numMap[prevCh]
		} else {
			num += numMap[b]
		}
		prevCh = b
	}

	return num
}
