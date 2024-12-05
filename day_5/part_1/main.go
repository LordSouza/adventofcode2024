package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type no struct {
	page        int
	less_then   []int
	bigger_then []int
}

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
	result := strings.Split(string(b), "\n\n")
	var order []string
	var updates []string
	order = strings.Split(result[0], "\n")
	updates = strings.Split(result[1], "\n")

	var order_int [][]int
	var updates_int [][]int

	for i := range order {
		j := strings.Split(order[i], "|")
		j_f, err := strconv.Atoi(j[0])
		var ord []int
		if err != nil {
			fmt.Println(err)
		}
		ord = append(ord, j_f)

		j_s, err := strconv.Atoi(j[1])
		if err != nil {
			fmt.Println(err)
		}
		ord = append(ord, j_s)

		order_int = append(order_int, ord)
	}
	for i := range updates {
		j := strings.Split(updates[i], ",")
		var ord []int
		for y := range j {
			x, err := strconv.Atoi(j[y])
			if err != nil {
				fmt.Println(err)
			}
			ord = append(ord, x)
		}
		updates_int = append(updates_int, ord)
	}
	var list_no []no
	m := make(map[int]int)
	for i := range order_int {
		for j := range order_int[i] {
			m[order_int[i][j]] = 0
		}
	}
	for i := range m {
		list_no = append(list_no, no{page: i})
	}
	for i := range order_int {
		for j := range list_no {
			if list_no[j].page == order_int[i][1] {
				list_no[j].bigger_then = append(list_no[j].bigger_then, order_int[i][0])
			}
			if list_no[j].page == order_int[i][0] {
				list_no[j].less_then = append(list_no[j].less_then, order_int[i][1])
			}
		}
	}

	for i := range updates_int {
		for x := range updates_int[i] {
			var n no
			for a := range list_no {
				if list_no[a].page == updates_int[i][x] {
					n = list_no[a]
				}
			}

		}
	}

	fmt.Println(list_no)
}

func check_bigger(noh no, x int) bool {
	for i := range noh.less_then {
		if i == x {
			return false
		}
	}
	return true
}
