package factorymethod

// 被封装的实际接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 工厂接口
type OperatorFactory interface {
	Create() Operator
}

// 接口实现基类
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// 加法实现类
type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

// 加法实现类工厂
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// 减法实现类
type MinusOperator struct {
	*OperatorBase
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

// 减法实现类工厂
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}
