package main

//func main() {
//	fmt.Println(checkPerfectNumber(28))
//}

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	total := 1

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			total += i
			if i*i < num {
				total += num / i
			}
		}
	}

	return num == total
}
