package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	result := strings.Split(string(b), "\n")
	var data [][]string
	for i := range result {
		y := strings.Split(result[i], "")

		data = append(data, y)
	}
	final := 0
	for i := range data {
		for j := range data[i] {
			if data[i][j] == "X" {
				final = final + check_words(data, i, j)
			}
		}
	}
	fmt.Println(final)
}
func check_words(data [][]string, i, j int) int {
	sum := 0
	xmas := []string{"X", "M", "A", "S"}
	// horizontal
	if j+len(xmas)-1 <= len(data[i])-1 && data[i][j] == xmas[0] && data[i][j+1] == xmas[1] && data[i][j+2] == xmas[2] && data[i][j+3] == xmas[3] {
		sum = sum + 1
	}
	// written backwards
	if j >= len(xmas)-1 && data[i][j] == xmas[0] && data[i][j-1] == xmas[1] && data[i][j-2] == xmas[2] && data[i][j-3] == xmas[3] {
		sum = sum + 1
	}
	// vertical
	if i+len(xmas)-1 <= len(data)-1 && data[i][j] == xmas[0] && data[i+1][j] == xmas[1] && data[i+2][j] == xmas[2] && data[i+3][j] == xmas[3] {
		sum = sum + 1
	}
	if i >= len(xmas)-1 && data[i][j] == xmas[0] && data[i-1][j] == xmas[1] && data[i-2][j] == xmas[2] && data[i-3][j] == xmas[3] {
		sum = sum + 1
	}
	// diagonal
	// top right
	if i >= len(xmas)-1 && j+len(xmas)-1 <= len(data[i])-1 && data[i][j] == xmas[0] && data[i-1][j+1] == xmas[1] && data[i-2][j+2] == xmas[2] && data[i-3][j+3] == xmas[3] {
		sum = sum + 1
	}
	// top left
	// fmt.Println(i, j, len(data[i]), len(data))
	if i >= len(xmas)-1 && j >= len(xmas)-1 && data[i][j] == xmas[0] && data[i-1][j-1] == xmas[1] && data[i-2][j-2] == xmas[2] && data[i-3][j-3] == xmas[3] {
		sum = sum + 1
	}
	// botton right
	if i+len(xmas)-1 <= len(data)-1 && j+len(xmas)-1 <= len(data[i])-1 && data[i][j] == xmas[0] && data[i+1][j+1] == xmas[1] && data[i+2][j+2] == xmas[2] && data[i+3][j+3] == xmas[3] {
		sum = sum + 1
	}
	// botton left
	if i+len(xmas)-1 <= len(data)-1 && j >= len(xmas)-1 && data[i][j] == xmas[0] && data[i+1][j-1] == xmas[1] && data[i+2][j-2] == xmas[2] && data[i+3][j-3] == xmas[3] {
		sum = sum + 1
	}

	return sum
}
