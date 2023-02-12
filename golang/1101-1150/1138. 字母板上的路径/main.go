package main

import (
	"fmt"
	"strings"
)

func alphabetBoardPath(target string) string {
	row, col := 0, 0
	var ans string

	for i := range target {
		cRow, cCol := getXAndY(target[i])
		rowStr := getRowString(row, cRow)
		colStr := getColString(col, cCol)

		if target[i] == 'z' {
			ans = fmt.Sprintf("%s%s%s!", ans, colStr, rowStr)
		} else {
			ans = fmt.Sprintf("%s%s%s!", ans, rowStr, colStr)
		}

		row, col = cRow, cCol
	}

	return ans
}

func getXAndY(c byte) (int, int) {
	v := int(c - 'a')
	return v / 5, v % 5
}

func getRowString(row, cRow int) string {
	rowPrint := "U"
	repeat := row - cRow
	if row < cRow {
		rowPrint = "D"
		repeat = -repeat
	}

	return strings.Repeat(rowPrint, repeat)
}

func getColString(col, cCol int) string {
	colPrint := "L"
	repeat := col - cCol
	if col < cCol {
		colPrint = "R"
		repeat = -repeat
	}

	return strings.Repeat(colPrint, repeat)
}
