package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(string(input))
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		input, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}

		if len(input) <= 10 {
			fmt.Println(string(input))
		} else {
			fmt.Println(string(input[:1]) + strconv.Itoa(len(input)-2) + string(input[len(input)-1]))
		}
	}
}
