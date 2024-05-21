package main

import (
	"strconv"
	"strings"
)

func ToCsvText(array [][]int) string {
	var rows []string

	for _, row := range array {
		var rowStr []string
		for _, num := range row {
			rowStr = append(rowStr, strconv.Itoa(num))
		}
		rows = append(rows, strings.Join(rowStr, ","))
	}

	return strings.Join(rows, "\n")
}
