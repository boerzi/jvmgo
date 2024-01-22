package jvmgo

import (
	"fmt"
	"leiyichen/jvmgo/classpath"
	"leiyichen/jvmgo/rtda/heap"
	"strings"
)

const VERSION = "0.0.1-toy"

type Cmd struct {
	Version   bool     `clop:"-v; --version"  usage:"print version and exit"`
	Classpath []string `clop:"-cp; --classpath"  usage:"classpath"`
	Jre       []string `clop:"-j; --jre; greedy" usage:"jre"`
}

func (c *Cmd) Main() {
	if c.Version {
		fmt.Println(VERSION)
	} else if len(c.Classpath) > 0 {
		if len(c.Classpath) < 2 {
			fmt.Println("classpath must be 2 param")
		} else {
			printUsage(c.Classpath)
		}
	} else if len(c.Jre) > 0 {
		fmt.Println(c.Jre[1])
		jre := c.Jre[0] //jre
		s := c.Jre[1]   //class
		s2 := c.Jre[2]  //class
		startJVM(jre, s2, s)
	}

}

func startJVM(reOption, cpOption string, class string) {
	cp := classpath.Parse(reOption, cpOption)
	fmt.Printf("classpath:%v class:%v \n", cp, class)

	classLoader := heap.NewClassLoader(cp)

	className := strings.Replace(class, ".", "/", -1)

	mainClass := classLoader.LoadClass(className) //class结构处理
	mainMethod := mainClass.GetMainMethod()       //先得到main方法（静态类 + main的特征）

	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", class)
	}
}

func printUsage(c []string) {
	fmt.Println(c)
	sprintf := fmt.Sprintf("classpath: %s class:%s args:%v\n", c[0], c[1], c[2:])
	fmt.Println(sprintf)
}
