package main

func main() {
	hammingWeight(00000000000000000000000000001011)
}

func hammingWeight(num uint32) int {
	ans := 0
	for num != 0 {
		if num%2 == 1 {
			ans++
		}
		num /= 2
	}

	return ans
}

func hammingWeightBetter(num uint32) int {
	sum := 0
	for num > 0 {
		sum++
		num = num & (num - 1)
	}
	return sum
}
