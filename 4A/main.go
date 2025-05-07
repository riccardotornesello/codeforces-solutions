package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	w, err := strconv.Atoi(string(input))
	if err != nil {
		log.Fatal(err)
	}

	if w > 2 && (w-2)&1 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
