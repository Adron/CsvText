package main

import "strconv"

func ToCsvText(array [][]int) string {
	var result string

	for i, row := range array {
		for j, num := range row {
			result += strconv.Itoa(num)
			if j < len(row)-1 {
				result += ","
			}
		}
		if i < len(array)-1 {
			result += "\n"
		}
	}

	return result
}
