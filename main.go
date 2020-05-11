package main

import (
	"fmt"

	"github.com/huangxinchun/go-demo/testpackage"
	"github.com/huangxinchun/hxcgo"
	"github.com/huangxinchun/hxcgo/my"
)

func main() {
	testpackage.TestNew()
	fmt.Println("main")
	my.TestMy()
	hxcgo.TestGin()

}
