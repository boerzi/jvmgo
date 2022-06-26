package main

import (
	"fmt"
	"github.com/guonaihong/clop"
	"leiyichen/jvmgo"
)

func main() {
	c := &jvmgo.Cmd{}
	err := clop.Bind(c)
	c.Main()
	fmt.Printf("%#v, %s\n", c, err)
}
