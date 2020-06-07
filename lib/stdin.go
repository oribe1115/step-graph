package lib

import (
	"bufio"
	"os"
)

var (
	stdin *bufio.Scanner
)

func InitStdin() {
	stdin = bufio.NewScanner(os.Stdin)
}

func ReadLine() string {
	stdin.Scan()
	return stdin.Text()
}
