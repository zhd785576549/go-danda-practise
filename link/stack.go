package link

// MaxStack 最大栈的数量
var MaxStack = 100

// Stack 栈
type Stack struct {
	Data []interface{}
	Top  int
}

// StackInit 栈初始化
func StackInit(s *Stack) error {
	if s.Data == nil {
		s.Data = make([]interface{}, MaxLinkLength)
		s.Top = -1
		return nil
	}

	return ErrAlreadyInit
}

// StackPush 入栈
func StackPush(s *Stack, data interface{}) error {
	if s.Top == MaxLinkLength {
		return ErrFull
	}
	s.Top++
	s.Data[s.Top] = data
	return nil
}

// StackPop 出栈
func StackPop(s *Stack, data interface{}) error {
	if s.Top == 0 {
		return ErrEmpty
	}

	data = s.Data[s.Top]
	s.Data[s.Top] = nil
	s.Top--
	return nil
}
