package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func Input(value string) string {
	fmt.Print(value)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	result := scanner.Text()
	return result
}
