package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct {
	action  string
	operand []float32
}

// __init__ sort of.
func NewCalculator(action string) *Calculator {
	var operand []float32
	calculator := Calculator{
		action:  action,
		operand: operand,
	}
	return &calculator
}

func (calc Calculator) calculate() float32 {
	var result float32
	switch calc.action {
	case "add":
		result = calc.add(calc.operand)
	case "subtract":
		result = calc.subtract(calc.operand)
	case "multiply":
		result = calc.mulitply(calc.operand)
	case "divide":
		result = calc.divide(calc.operand)
	}
	return result
}

func (calc Calculator) add(values []float32) float32 {
	var result float32
	for _, value := range values {
		result += float32(value)
	}
	return result
}

func (calc Calculator) subtract(values []float32) float32 {
	var result float32
	result = values[0]
	for i := 1; i < len(values); i++ {
		result = result - values[i]
	}
	return result
}

func (calc Calculator) divide(values []float32) float32 {
	var result float32
	result = values[0]
	for i := 1; i < len(values); i++ {
		result /= values[i]
	}
	return result
}

func (calc Calculator) mulitply(values []float32) float32 {
	var result float32
	result = values[0]
	for i := 1; i < len(values); i++ {
		result *= values[i]
	}
	return result
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi! This is calculatorProMax. You can add, subtract, multiply & divide. Enter stop anytime to quit!")
	var operation string

	operation = "start"
	for operation != "stop" {
		fmt.Print("Choose an operation (+, -, *, /) : ")
		input, _ := reader.ReadString('\n')
		trimmedInput := strings.TrimSpace(input)

		for trimmedInput != "+" && trimmedInput != "-" && trimmedInput != "/" && trimmedInput != "*" {
			if trimmedInput == "stop" {
				os.Exit(0)
			}
			fmt.Print("Invalid operation. please choose one of :(+, -, *, /) : ")
			input, _ := reader.ReadString('\n')
			trimmedInput = strings.TrimSpace(input)
		}

		switch trimmedInput {
		case "+":
			operation = "add"
		case "-":
			operation = "subtract"
		case "/":
			operation = "divide"
		case "*":
			operation = "multiply"
		case "stop":
			os.Exit(0)
		}

		fmt.Printf("hmm, i see you wanna %v.\n", operation)
		calculator := NewCalculator(operation)

		fmt.Printf("How many numbers would you like to %v : ", operation)
		input_, _ := reader.ReadString('\n')
		iterate, err := strconv.ParseInt(strings.TrimSpace(input_), 10, 64)
		if strings.TrimSpace(input_) == "stop" {
			os.Exit(0)
		}
		for err != nil {
			fmt.Printf("Invalid input. How many numbers would you like to %v : ", operation)
			input_, _ := reader.ReadString('\n')
			iterate, err = strconv.ParseInt(strings.TrimSpace(input_), 10, 64)
			if strings.TrimSpace(input_) == "stop" {
				os.Exit(0)
			}
		}

		for i := 0; i < int(iterate); i++ {
			fmt.Printf("Enter input %d : ", i+1)
			input_, _ = reader.ReadString('\n')
			if strings.TrimSpace(input_) == "stop" {
				os.Exit(0)
			}
			aFloat, err_ := strconv.ParseFloat(strings.TrimSpace(input_), 64)

			for err_ != nil {
				fmt.Printf("Invalid input. Re-Enter input %d : ", i+1)
				input_, _ := reader.ReadString('\n')
				aFloat, err_ = strconv.ParseFloat(strings.TrimSpace(input_), 64)
			}
			calculator.operand = append(calculator.operand, float32(aFloat))
		}
		result := calculator.calculate()
		fmt.Println("result is : ", result)

	}
}
