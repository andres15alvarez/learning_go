package utils

import (
	"fmt"
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

// This function is a calculator that receive two integer numbers and an arithmetic operation
// like +,-,*,/; and print in the console the result
func Calculate() {
	firstOperator := Input("Enter the first number: ")
	secondOperator := Input("Enter the second number: ")
	operator := Input("Enter the operation (+,-,*,/): ")
	firstValue, err := strconv.ParseInt(firstOperator, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	secondValue, err := strconv.ParseInt(secondOperator, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	calc := Calculator{int(firstValue), int(secondValue), operator, 0}
	calc.handleCalculation()
	fmt.Printf("Result: %.2f\n", calc.result)
}
