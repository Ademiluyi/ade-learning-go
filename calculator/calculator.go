package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome! This is a simple calculator app & you can add up to 9 values at once!")
	var sum int
	fmt.Print("\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter how many numbers you would like to add  : ")
	input, _ := reader.ReadString('\n')
	anInt, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err == nil && int(anInt) >= 0 && int(anInt) <= 9 {
		for i := 0; i < int(anInt); i++ {
			fmt.Printf("Enter value %d : ", i+1)
			input, _ := reader.ReadString('\n')
			anInt, err1 := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
			for err1 != nil {
				fmt.Printf("Invalid input. Re-Enter value %d", i+1)
				input, _ := reader.ReadString('\n')
				anInt, err1 = strconv.ParseInt(strings.TrimSpace(input), 10, 64)
			}
			sum += int(anInt)
		}
	} else {
		panic("Error!! Invalid input!")
	}
	fmt.Println("Total sum : ", sum)
}
