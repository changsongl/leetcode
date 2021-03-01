package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([][]byte, numRows)
	down := true
	i := 0

	for _, letter := range s {
		rows[i] = append(rows[i], byte(letter))
		if down {
			i++
		} else {
			i--
		}

		if i == 0 {
			down = true
		} else if i == numRows-1 {
			down = false
		}
	}

	ans := ""
	for _, row := range rows {
		ans += string(row)
	}
	return ans
}
