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

	var data [][]int

	for i := range result {
		y := strings.Split(result[i], " ")
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
		decre := make([]int, len(data[i]))
		incre := make([]int, len(data[i]))

		copy(decre, data[i])
		copy(incre, data[i])
		if check_decrease(decre) || check_increase(incre) {
			resposta = resposta + 1
		}
	}
	fmt.Println(resposta)

}

func check_decrease(values []int) bool {
	first_decrease := make([]int, len(values))
	second_decrease := make([]int, len(values))
	level_dump := false
	for i := range values {
		if i == len(values)-1 {
			break
		}

		if values[i] < values[i+1] || values[i] == values[i+1] || bigger_thar_three(values[i], values[i+1]) {
			first_decrease = remove(values, i+1)
			second_decrease = remove(values, i)

			level_dump = !level_dump
			break
		}

	}

	if level_dump {
		if decrease_for(first_decrease) || decrease_for(second_decrease) {
			return true
		} else {
			return false
		}
	}
	return true
}
func check_increase(values []int) bool {
	first_increase := make([]int, len(values))
	second_increase := make([]int, len(values))
	level_dump := false

	for i := range values {
		if i == len(values)-1 {
			break
		}

		if values[i] > values[i+1] || values[i] == values[i+1] || bigger_thar_three(values[i], values[i+1]) {
			first_increase = remove(values, i+1)
			second_increase = remove(values, i)
			level_dump = !level_dump
			break
		}

	}

	if level_dump {
		if increase_for(first_increase) || increase_for(second_increase) {
			return true
		} else {
			return false
		}
	}

	return true
}

func decrease_for(slice []int) bool {
	for i := range slice {
		if i == len(slice)-1 {
			break
		}

		if slice[i] < slice[i+1] || slice[i] == slice[i+1] || bigger_thar_three(slice[i], slice[i+1]) {
			return false
		}
	}
	return true
}
func increase_for(slice []int) bool {
	for i := range slice {
		if i == len(slice)-1 {
			break
		}

		if slice[i] > slice[i+1] || slice[i] == slice[i+1] || bigger_thar_three(slice[i], slice[i+1]) {
			return false
		}
	}
	return true
}
func remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func bigger_thar_three(a, b int) bool {
	return math.Abs(math.Abs(float64(a))-math.Abs(float64(b))) > 3
}
