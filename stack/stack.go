/**
栈（链表实现）
 */
package stack

import "fmt"

type StackItem struct {
	Item interface{}
	Next *StackItem
}

type Stack struct {
	Sp    *StackItem
	Depth int
}

func New() *Stack {
	return new(Stack)
}

func (s *Stack)IsEmpty() bool {
	return s.Depth == 0
}

func (s *Stack) Push(item interface{}) {
	s.Sp = &StackItem{Item: item, Next: s.Sp}
	s.Depth++
}

func (s *Stack) Pop() interface{} {
	if s.Depth > 0 {
		item := s.Sp.Item
		s.Sp = s.Sp.Next
		s.Depth--
		return item

	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.Depth > 0 {
		return s.Sp.Item
	}
	return nil
}

func (s *Stack) Print() {
	sp := s.Sp
	for sp != nil {
		fmt.Println(sp.Item)
		sp = sp.Next
	}
}

//判断括号表达式是否正确
func IsBalanceParentHesis(str string) bool {
	stack := new(Stack)
	for _, ch := range str {
		switch ch {
		case '{','[','(':
			stack.Push(ch)
		case '}':
			c := stack.Pop()
			if c != '{' {
				return false
			}
		case ']':
			c := stack.Pop()
			if c != '[' {
				return false
			}
		case ')':
			c := stack.Pop()
			if c != '(' {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

//中缀表达式转后缀表达式
func InfixToPostfix(expn string) string {
	stk := new(Stack)
	output := ""
	for _,ch := range expn {
		if ch >='0' && ch <= '9' {
			output = output + string(ch)
		} else {
			switch ch {
			case '+','-','*','/','%':
				for !stk.IsEmpty() && precedence(ch) < precedence(stk.Peek().(rune)) {
					out := stk.Pop().(rune)
					output = output + " " + string(out)
				}
				stk.Push(ch)
				output = output + " "
			case '(':
				stk.Push(ch)
			case ')':
				for !stk.IsEmpty() && stk.Peek() != '(' {
					out := stk.Pop().(rune)
					output = output + " " + string(out)
				}
				if !stk.IsEmpty() && stk.Peek() == '(' {
					stk.Pop()
				}
			}
		}
	}
	for !stk.IsEmpty() {
		out := stk.Pop().(rune)
		output = output + " " + string(out)
	}
	return output
}

func precedence(x rune) int {
	if x == '(' {
		return 0
	}
	if x == '+' || x == '-' {
		return 1
	}
	if x == '*' || x == '/' || x == '%' {
		return 2
	}
	if x == '^' {
		return 3
	}
	return 4
}