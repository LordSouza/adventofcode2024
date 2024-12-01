package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
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

	var ids_1 []int
	var ids_2 []int

	result := strings.Split(string(b), "\n")

	for i := range result {
		y := strings.Split(result[i], "   ")

		if len(y) == 2 {
			first, err := strconv.Atoi(y[0])
			if err != nil {
				log.Fatal(err)
			}

			ids_1 = append(ids_1, first)
			second, err := strconv.Atoi(y[1])
			if err != nil {
				log.Fatal(err)
			}
			ids_2 = append(ids_2, second)
		}
	}
	slices.Sort(ids_1)
	slices.Sort(ids_2)
	var resposta int
	for i := range ids_1 {
		resposta += int(math.Abs(float64(ids_2[i] - ids_1[i])))
		fmt.Println(resposta, i, ids_2[i], ids_1[i], ids_2[i]-ids_1[i])
	}

}
