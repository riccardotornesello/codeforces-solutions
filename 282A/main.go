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

	x := 0

	for i := 0; i < n; i++ {
		input, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}

		inputStr := string(input)

		// Not good looking but it works efficiently for this case
		if inputStr == "X++" || inputStr == "++X" {
			x += 1
		} else if inputStr == "X--" || inputStr == "--X" {
			x -= 1
		}
	}

	fmt.Println(x)
}
