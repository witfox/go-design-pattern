package method

/*工厂方法
* 每个类都有自己的工厂类
 */

// 运算实现类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 运算工厂接口
type OperatorFactory interface {
	Create() Operator
}

// 实现类共用部分
type OperatorBase struct {
	a, b int
}

func (ob *OperatorBase) SetA(a int) {
	ob.a = a
}

func (ob *OperatorBase) SetB(b int) {
	ob.b = b
}

// 加法类
type PlusOperator struct {
	*OperatorBase
}

func (po PlusOperator) Result() int {
	return po.a + po.b
}

// 加法类工厂
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// 减法类
type MinusOperator struct {
	*OperatorBase
}

func (mo MinusOperator) Result() int {
	return mo.a - mo.b
}

// 减法类工厂
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}
