package method

import (
	"fmt"
	"testing"
)

func TestMethod(t *testing.T) {
	factory := PlusOperatorFactory{}
	op := factory.Create()
	op.SetA(2)
	op.SetB(3)
	fmt.Println(op.Result())

	factory2 := MinusOperatorFactory{}
	op2 := factory2.Create()
	op2.SetA(2)
	op2.SetB(3)
	fmt.Println(op2.Result())
}
