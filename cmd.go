package jvmgo

import (
	"fmt"
	"leiyichen/jvmgo/classpath"
	"leiyichen/jvmgo/rtda/heap"
	"strconv"
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
		jre := c.Jre[0]                        //jre地址
		classPath := c.Jre[1]                  //类路径
		className := c.Jre[2]                  //class
		flag, _ := strconv.ParseBool(c.Jre[3]) //flag
		startJVM(jre, className, classPath, flag, nil)
	}

}

func startJVM(reOption, cpOption string, class string, flag bool, args []string) {
	cp := classpath.Parse(reOption, cpOption)
	fmt.Printf("classpath:%v \n", cp)
	fmt.Printf("class:%v \n", class)

	classLoader := heap.NewClassLoader(cp, flag) //初始化类加载器

	className := strings.Replace(class, ".", "/", -1)
	fmt.Println("className:" + className)

	mainClass := classLoader.LoadClass(className) //class结构处理
	mainMethod := mainClass.GetMainMethod()       //先得到main方法（静态类 + main的特征）

	if mainMethod != nil {
		interpret(mainMethod, flag, args) //开始执行
	} else {
		fmt.Printf("Main method not found in class %s\n", class)
	}
}

func printUsage(c []string) {
	fmt.Println(c)
	sprintf := fmt.Sprintf("classpath: %s class:%s args:%v\n", c[0], c[1], c[2:])
	fmt.Println(sprintf)
}
