package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())

	processName := os.Args[0]
	fmt.Println(os.Args)
	fmt.Println(processName)
}
