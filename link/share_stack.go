package link

// @Title share_stack
// @Description 共享栈

// MaxShareStackLength 共享栈最大数量
var MaxShareStackLength int = 100

// StackType 共享栈类型
type StackType int

const (
	// STypeTop1 top1 的栈
	STypeTop1 StackType = 1

	// STypeTop2 top2 的栈
	STypeTop2 StackType = 2
)

// ShareStack 共享栈数据结构
type ShareStack struct {
	Data []interface{}
	Top1 int
	Top2 int
}

// ShareStackInit 初始化共享栈
func ShareStackInit(s *ShareStack) error {
	if s.Data != nil {
		return ErrAlreadyInit
	}
	// | xxx1 | xxx1 | ...| xxx2 | xxx2|
	// |   0  |   1  | ...|   -2 |  -1 |
	// 栈1从左向右存放
	// 栈2从右向左存放
	// 0, 1, -1, -2 表示下标，0,1是正向下标，-1,-2是逆向下标，-1代表最后一个元素

	s.Data = make([]interface{}, MaxShareStackLength)
	s.Top1 = 0
	s.Top2 = MaxShareStackLength - 1

	return nil
}

// ShareStackPush 共享栈压栈
func ShareStackPush(s *ShareStack, t StackType, data interface{}) error {
	// | ..... | xxx1 | xxx2 | .....|
	// 当满足条件 top2 - top1 = 1时，或者 top2 = top1 + 1，栈是满的
	//
	// | xxx2 | xxx2 | ... |
	// 当top2 = 0时，栈是满的，都是栈2的数据
	//
	// | ... | xxx1 | xxx1 |
	// 当top1 = MaxShareStackLength-1时，栈是满的，都是栈1的数据
	//
	// 以上总结 top2 - top1 = 1，表示栈是满的
	if s.Top2-s.Top1 == 1 {
		return ErrFull
	}

	if t == STypeTop1 {
		s.Top1++
		s.Data[s.Top1] = data
	} else if t == STypeTop2 {
		s.Top2--
		s.Data[s.Top2] = data
	}

	return nil
}

// ShareStackPop 共享栈出栈
func ShareStackPop(s *ShareStack, t StackType, data interface{}) error {
	if (s.Top1 == 0 && t == STypeTop1) || (s.Top2 == MaxShareStackLength-1 && t == STypeTop2) {
		return ErrEmpty
	}

	if t == STypeTop1 {
		s.Top1--
		data = s.Data[s.Top1]
	} else if t == STypeTop2 {
		s.Top2++
		data = s.Data[s.Top2]
	}

	return nil
}
