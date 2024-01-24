package jvmgo

import (
	"fmt"
	"leiyichen/jvmgo/instructions"
	"leiyichen/jvmgo/instructions/base"
	"leiyichen/jvmgo/rtda"
	"leiyichen/jvmgo/rtda/heap"
)

func interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewThread()       //初始化线程
	frame := thread.NewFrame(method) //帧初始化获取代码
	thread.PushFrame(frame)          //压帧

	defer catchErr(thread)
	loop(thread, logInst)
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}

	for {
		frame := thread.CurrentFrame() //获取当前方法区（帧）
		pc := frame.NextPC()           //得到帧对应的程序计数器
		thread.SetPC(pc)

		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()                //读取指令
		inst := instructions.NewInstruction(opcode) //匹配指令
		inst.FetchOperands(reader)                  //取操作数
		frame.SetNextPC(reader.PC())                //得到下一个程序计数器

		if logInst {
			logInstruction(frame, inst)
		}

		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
