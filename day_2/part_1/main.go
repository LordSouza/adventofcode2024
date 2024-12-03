package main

import (
	"fmt"
	"io"
	"log"

	"math"
	"os"
	"strconv"
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

	result := strings.Split(string(b), "\n")

	var data [][] int

	for i := range result {
		y:= strings.Split(result[i], " ")
		var y_int []int
		for j := range y {
			v, err := strconv.Atoi(y[j])
			if err != nil {
				fmt.Println(err)
			}
			y_int = append(y_int, v)
		}
		data = append(data, y_int)
	}

	resposta := 0
	// checagens
	for i := range data {
		
		if check_decrease(data[i]) || check_increase(data[i]) {
			resposta = resposta + 1
		}
	}
	fmt.Println(resposta)

}

func check_decrease(values [] int) bool {
	for i := range values {
		if i == len(values) - 1 {

			break
		}
		
		if values[i] > values[i + 1] || values[i] == values[i + 1] || math.Abs(float64(values[i] - values[i + 1])) > 3 {
			return false
		}

	}
	return true
}
func check_increase(values [] int) bool {
	for i := range values {

		if i == len(values) - 1 {
			break
		}
		
		if values[i] < values[i + 1] || values[i] == values[i + 1] || math.Abs(float64(values[i] - values[i + 1])) > 3 {
			return false
		}

	}
	return true
}
