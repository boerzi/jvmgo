package base

import "leiyichen/jvmgo/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) //获取操作数
	Execute(frame *rtda.Frame)            //执行操作码
}

// NoOperandsInstruction 没有操作数的指令
type NoOperandsInstruction struct {
}

func (n *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

// BranchInstruction 跳转指令
type BranchInstruction struct {
	Offset int //偏移量
}

func (b *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	b.Offset = int(reader.ReadInt16())
}

// Index8Instruction 存储加载类指令根据索引存储局部变量表 索引由单字节操作给出
type Index8Instruction struct {
	Index uint //局部变量表的索引
}

func (i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint8())
}

// Index16Instruction 两字节输出
type Index16Instruction struct {
	Index uint //局部变量表的索引
}

func (i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i.Index = uint(reader.ReadUint16())

}
