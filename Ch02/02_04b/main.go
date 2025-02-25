package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(a int, b int) int {
	return a + b
}

func main() {
	var int1, int2 int
	var err error

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input first number: ")
	strFirstNumber, err := reader.ReadString('\n')
	int1, err = strconv.Atoi(strings.TrimSpace(strFirstNumber))
	if err != nil {
		fmt.Printf("%T, %v, \n", int1, int1)
		return
	} else {
		fmt.Println("Input second number: ")
		strSecondNumber, err := reader.ReadString('\n')
		int2, err = strconv.Atoi(strings.TrimSpace(strSecondNumber))
		if err != nil {
			fmt.Printf("%T, %v, \n", int2, int2)
			return
		}
	}
	fmt.Println("Sum of the two numbers: ", sum(int1, int2))
}
