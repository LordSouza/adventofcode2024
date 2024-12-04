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
			if data[i][j] == "A" {
				final = final + check_words(data, i, j)
			}
		}
	}
	fmt.Println(final)
}
func check_words(data [][]string, i, j int) int {
	sum := 0
	xmas := []string{"M", "S"}
	// check x
	if i-1 >= 0 && i+1 <= len(data)-1 && j-1 >= 0 && j+1 <= len(data[i])-1 {
		//  M . M
		//  . A .
		//  S . S
		if data[i-1][j-1] == xmas[0] && data[i+1][j+1] == xmas[1] && data[i-1][j+1] == xmas[0] && data[i+1][j-1] == xmas[1] {
			sum = sum + 1
		}
		//  M . S
		//  . A .
		//  M . S
		if data[i-1][j-1] == xmas[0] && data[i+1][j+1] == xmas[1] && data[i-1][j+1] == xmas[1] && data[i+1][j-1] == xmas[0] {
			sum = sum + 1
		}
		//  S . S
		//  . A .
		//  M . M
		if data[i-1][j-1] == xmas[1] && data[i+1][j+1] == xmas[0] && data[i-1][j+1] == xmas[0] && data[i+1][j-1] == xmas[1] {
			sum = sum + 1
		}
		//  S . M
		//  . A .
		//  S . M
		if data[i-1][j-1] == xmas[1] && data[i+1][j+1] == xmas[0] && data[i-1][j+1] == xmas[1] && data[i+1][j-1] == xmas[0] {
			sum = sum + 1
		}
	}
	return sum
}
