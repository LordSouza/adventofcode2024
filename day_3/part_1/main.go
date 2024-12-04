package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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
	r, _ := regexp.Compile("mul\\([0-9]+\\,[0-9]+\\)")
	result := r.FindAllString(string(b), -1)
	data := 0
	for i := range result {
		x := strings.Split(result[i][4:], ",")
		x[1] = TrimSuffix(x[1], ")")
		x_int, err := strconv.Atoi(x[0])
		if err != nil {
			fmt.Println(err)
		}
		y_int, err := strconv.Atoi(x[1])
		if err != nil {
			fmt.Println(err)
		}
		data = data + (x_int * y_int)
	}
	fmt.Println(data)
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
