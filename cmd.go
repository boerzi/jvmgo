package jvmgo

import (
	"fmt"
	"leiyichen/jvmgo/classpath"
	"strings"
)

type Cmd struct {
	Version   bool     `clop:"-v; --version"  usage:"number nonempty output lines, overrides"`
	Classpath []string `clop:"-cp;greedy; --classpath"  usage:"string" usage:"number nonempty output lines, overrides"`
	Jre       []string `clop:"-jre;greedy"  usage:"string" usage:"number nonempty output lines, overrides"`
}

type ParseCompleteCmdLineParam struct {
	Classpath string
	Class     string
	args      []string
}

func (c *Cmd) Main() {
	if c.Version {
		fmt.Println("version 0.0.1")
	} else if len(c.Classpath) > 0 {
		if len(c.Classpath) < 2 {
			fmt.Println("classpath must be 2 param")
		} else {
			printUsage(c.Classpath)
		}
	} else if len(c.Jre) > 0 {
		jre := c.Jre[0]
		s := c.Jre[1]
		startJVM(jre, s)
	}

}

func startJVM(reOption, class string) {
	cp := classpath.Parse(reOption, "")
	fmt.Printf("classpath:%v class:%v", cp, class)
	className := strings.Replace(class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", class)
		return
	}

	fmt.Printf("class data:%v\n", classData)

}

func printUsage(c []string) {
	fmt.Println(c)
	sprintf := fmt.Sprintf("classpath: %s class:%s args:%v\n", c[0], c[1], c[2:])
	fmt.Println(sprintf)
}
