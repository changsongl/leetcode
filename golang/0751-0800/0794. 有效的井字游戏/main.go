package main

import "fmt"

func main() {
	//fmt.Println(validTicTacToe([]string{"XXX", "   ", "OOO"}), false)
	//fmt.Println(validTicTacToe([]string{"XOX", " X ", "   "}), false)
	//fmt.Println(validTicTacToe([]string{"O  ", "   ", "   "}), false)
	//fmt.Println(validTicTacToe([]string{"XOX", "O O", "XOX"}), true)
	//
	//fmt.Println(validTicTacToe([]string{"XXX", "OOX", "OOX"}), true)

	fmt.Println(validTicTacToe([]string{"XXO", "XOX", "OXO"}), false)
	//"XXO"
	//"XOX"
	//"OXO"
}

// 支持大于3x3
func validTicTacToe(board []string) bool {
	height := len(board)
	width := len(board[0])

	countX := 0
	countO := 0
	var winX, winO [][][]int

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == ' ' {
				continue
			}

			var wins [][][]int
			incs := [][]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}}
			for _, inc := range incs {
				if (i-inc[0] >= 0 && i-inc[0] < height && j-inc[1] >= 0 && j-inc[1] < width) && board[i-inc[0]][j-inc[1]] == board[i][j] {
					continue
				}

				checkI, checkJ := i, j
				checkWins := [][]int{{checkI, checkJ}}
				for (checkI+inc[0] < height && checkI+inc[0] >= 0 && checkJ+inc[1] < width && checkJ+inc[1] >= 0) &&
					board[checkI+inc[0]][checkJ+inc[1]] == board[checkI][checkJ] {
					checkWins = append(checkWins, []int{checkI + inc[0], checkJ + inc[1]})
					checkI = checkI + inc[0]
					checkJ = checkJ + inc[1]
				}

				if len(checkWins) > 5 {
					return false
				} else if len(checkWins) >= 3 {
					wins = append(wins, checkWins)
				}
			}

			if board[i][j] == 'X' {
				countX++
				winX = append(winX, wins...)
			} else {
				countO++
				winO = append(winO, wins...)
			}
		}
	}

	if countX-countO != 0 && countX-countO != 1 {
		return false
	} else if len(winX) != 0 && len(winO) != 0 {
		return false
	} else if len(winX) > 0 && countX-countO == 0 {
		return false
	} else if len(winO) > 0 && countX-countO == 1 {
		return false
	}

	checkWins := append(winX, winO...)
	if len(checkWins) <= 1 {
		return true
	}

	crossI, crossJ := -1, -1
	for i := 1; i < len(checkWins); i++ {
		for checkIndex, check := range checkWins[i] {
			for checkIndex2, check2 := range checkWins[0] {
				if check[0] == check2[0] && check[1] == check2[1] {
					if crossJ != -1 && (crossJ != check2[1] || crossI != check2[0]) {
						return false
					} else if crossJ == -1 {
						crossI = check2[0]
						crossJ = check2[1]
					}

					if checkIndex-3 >= 0 || checkIndex+3 < len(checkWins[i]) {
						return false
					} else if checkIndex2-3 >= 0 || checkIndex2+3 < len(checkWins[0]) {
						return false
					}
				}
			}
		}
	}

	return true
}
