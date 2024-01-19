package heap

// SymRef 符号引用所在的运行时常量池指针
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}
