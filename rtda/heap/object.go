package heap

type Object struct {
	class  *Class
	fields Slots //存放实例变量
}

func (s *Object) Class() *Class {
	return s.class
}
func (s *Object) Fields() Slots {
	return s.fields
}

func (s *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(s.class)
}
