package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Calculator struct {
	firstOperator  int
	secondOperator int
	operation      string
	result         float64
}

func (c *Calculator) add() {
	c.result = float64(c.firstOperator) + float64(c.secondOperator)
}

func (c *Calculator) sub() {
	c.result = float64(c.firstOperator) - float64(c.secondOperator)
}

func (c *Calculator) mul() {
	c.result = float64(c.firstOperator) * float64(c.secondOperator)
}

func (c *Calculator) div() {
	c.result = float64(c.firstOperator) / float64(c.secondOperator)
}

func (c *Calculator) handleCalculation() {
	switch c.operation {
	case "+":
		c.add()
		break

	case "-":
		c.sub()
		break

	case "*":
		c.mul()
		break

	case "/":
		c.div()
		break

	default:
		fmt.Println("Operation not allowed.")
		c.result = 0
		break
	}
}

func enterValue(message string) string {
	fmt.Print(message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operator := scanner.Text()
	return operator
}

func convertStrToInt(value string) (int64, error) {
	number, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, errors.New("Cannot convert to int")
	}
	return number, nil
}

// This function is a calculator
func Calculate() {
	firstOperator := enterValue("Enter the first number: ")
	secondOperator := enterValue("Enter the second number: ")
	operator := enterValue("Enter the operation (+,-,*,/): ")
	firstValue, err := convertStrToInt(firstOperator)
	if err != nil {
		fmt.Println(err)
	}
	secondValue, err := convertStrToInt(secondOperator)
	if err != nil {
		fmt.Println(err)
	}
	calc := Calculator{int(firstValue), int(secondValue), operator, 0}
	calc.handleCalculation()
	fmt.Printf("Result: %.2f\n", calc.result)
}

func Addition() {
	fmt.Println(2 + 2)
}
