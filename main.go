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
	result         int
}

func (c *Calculator) add() {
	c.result = c.firstOperator + c.secondOperator
}

func (c *Calculator) sub() {
	c.result = c.firstOperator - c.secondOperator
}

func (c *Calculator) mul() {
	c.result = c.firstOperator * c.secondOperator
}

func (c *Calculator) div() {
	c.result = c.firstOperator / c.secondOperator
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

func main() {
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
	fmt.Printf("Result: %d\n", calc.result)
}
