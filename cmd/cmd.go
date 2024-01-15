package main

import (
	"github.com/guonaihong/clop"
	"leiyichen/jvmgo"
)

func main() {
	c := &jvmgo.Cmd{}
	err := clop.Bind(c)
	if err != nil {
		panic(err)
	}
	c.Main()
}
