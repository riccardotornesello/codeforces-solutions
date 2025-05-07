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

	validSolutions := 0

	for i := 0; i < n; i++ {
		input, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}

		sureMembers := 0
		if input[0] == '1' {
			sureMembers += 1
		}
		if input[2] == '1' {
			sureMembers += 1
		}
		if input[4] == '1' {
			sureMembers += 1
		}

		if sureMembers >= 2 {
			validSolutions += 1
		}
	}

	fmt.Println(validSolutions)
}
