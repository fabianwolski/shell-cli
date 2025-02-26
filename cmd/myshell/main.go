package main

import (
	"bufio"
	"fmt"
	"os"
)

var _ = fmt.Fprint

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	// Wait for user input
	bufio.NewReader(os.Stdin).ReadString('\n')
}
