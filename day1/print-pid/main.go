package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	//获取进程id
	fmt.Println(os.Getpid())
	//获取进程父id
	fmt.Println(os.Getppid())
	//进程名称
	processName := os.Args[0]
	fmt.Println(os.Args)
	fmt.Println(processName)

	//获取实体参数

	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	//获得进程Flag
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("number", 42, "a int")

	flag.Parse()
	fmt.Println(*wordPtr)
	fmt.Println(*numbPtr)
}
