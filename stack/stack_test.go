package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	var stack = new(Stack)
	var max = 1000
	for i:=0;i<=max;i++ {
		stack.Push(i)
	}
	for i:=max;i>=0;i-- {
		item := stack.Pop()
		if item != i {
			t.Error("TestStack failed...")
		}
	}
}

func TestIsBalanceParentHesis(t *testing.T) {
	str := "{[(4,4,4)]}"
	res := IsBalanceParentHesis(str)
	if !res {
		t.Error("TestIsBalanceParentHesis failed...")
	}
	str = "{[(()]}"
	res = IsBalanceParentHesis(str)
	if res {
		t.Error("TestIsBalanceParentHesis failed...")
	}
}

func TestInfixToPostfix(t *testing.T) {
	expn := "(1+3)*4"
	value := InfixToPostfix(expn)
	if value != "1 3 + 4 *" {
		t.Error("TestInfixToPostfix failed...")
	}
	expn = "10+((3))*5/(16-4)"
	value = InfixToPostfix(expn)
	if value != "10 3 5 16 4 - / * +" {
		t.Error("TestInfixToPostfix failed...")
	}
}