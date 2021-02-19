package link

// @Title link 链表包
// @Description 单项链表数据结构练习，链表空时是nil，只有Next一个方向进行查询，每个节点都可以携带任意对象

// Node 节点
type Node struct {
	Data interface{} // 任意对象的指针
	Next *Node       // 下一个节点
}

// Pos 位置
type Pos int

const (
	// Head 头部
	Head Pos = 1

	// Middle 中间
	Middle Pos = 2

	// Tail 尾部
	Tail Pos = 3
)

// Link 连表头
type Link *Node

// Init 初始化
func Init(l *Link) error {
	*l = nil
	return nil
}

// AddNode 添加节点
// @title AddNode
// @description 添加节点到给定的链表中
// @param l  *Link 链表头
// @param pos Pos 添加的位置，Head头部，Tail尾部，Middle中间
// @param node *Node 添加的节点
// @param step int 当Pos被设置成Middle时，step表示插入那个位置
func AddNode(l *Link, pos Pos, node *Node, step int) error {

	if l == nil && pos != Head {
		return ErrEmptyLink
	}

	switch pos {
	case Head:
		return addToHead(l, node)
	case Middle:
		return addToMiddle(l, node, step)
	case Tail:
		return addToTail(l, node)
	default:
		return ErrUnkownPosition
	}
}

// addToHead 添加到头部
func addToHead(l *Link, node *Node) error {
	head := *l
	if head == nil {
		*l = node
	} else {
		node.Next = head
		*l = node
	}
	return nil
}

// addToTail 添加到尾部
func addToTail(l *Link, node *Node) error {
	head := *l
	if head == nil {
		*l = node
	} else {
		tmp := head
		for tmp.Next != nil {
			tmp = tmp.Next
		}

		tmp.Next = node
	}

	return nil
}

// addToMiddle 添加节点到step的节点后面
func addToMiddle(l *Link, node *Node, step int) error {
	head := *l
	tmp := head
	curStep := 0

	for curStep < step-1 {
		if tmp == nil {
			return ErrOverflow
		}

		tmp = tmp.Next
		curStep++
	}

	// 当前节点可能是空
	if tmp == nil {
		tmp.Next = node
	} else {
		node.Next = tmp.Next
		tmp.Next = node
	}

	return nil
}

// GetCount 获取链表总数
func GetCount(l Link) (int, error) {
	count := 0
	tmp := l

	for {

		if tmp == nil {
			break
		}
		count++
		tmp = tmp.Next
	}

	return count, nil
}

// GetNodeByPosition 通过位置获取节点
func GetNodeByPosition(l Link, pos int, node **Node) error {
	tmp := l

	// 判断空列表
	if l == nil {
		return ErrEmptyLink
	}

	// 找出位置为pos - 1的节点
	for i := 0; i < pos-1; i++ {
		if tmp == nil {
			return ErrOverflow
		}
		tmp = tmp.Next
	}

	// 返回该节点的内容
	*node = tmp
	return nil
}

// Reverse 链表反转
func Reverse(l *Link) error {
	head := *l

	if head == nil || head.Next == nil {
		return nil
	}

	p1 := head
	p2 := head.Next

	//    p1      p2      p3
	// 1. 1 ----> 2 ----> 3 -----> 4 ... ---> nil
	//    p1      p2      p3
	// 2. 1 ----> 2       3 -----> 4 ... ---> nil
	//     <-----
	//           p1       p2       p3
	// 3. 1 ----> 2       3------> 4 ... ---> nil
	//     <-----
	//           p1       p2       p3
	// 4. 1-----> 2<----- 3        4 ... ---> nil
	//     <-----  -
	// ...
	// 5. 1-----> 2<------3<-------4 ...<----last
	//     <-----
	// 6. nil <---1<----2<------3<-------4 ...<----last
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}

	head.Next = nil
	*l = p1
	return nil
}

// DeleteNode 删除指定步长的节点
func DeleteNode(l *Link, step int) error {
	head := *l
	tmp := head

	if head == nil {
		return ErrEmptyLink
	}

	curStep := 0
	prev := tmp
	for curStep < step-1 {
		if tmp == nil {
			return ErrOverflow
		}

		prev = tmp
		tmp = tmp.Next
		curStep++
	}

	prev.Next = tmp.Next

	return nil
}

// LinkGenerator 迭代生成器，循环遍历链表并且通过channel返回节点内容
func LinkGenerator(l Link) chan Node {
	c := make(chan Node)
	// 这里是匿名方法的协程，主要作用是将连表的node写入channel中
	// 调用者使用channel不断读出channel中的数据
	go func() {
		tmp := l
		for {
			if tmp == nil {
				break
			}

			c <- (*tmp)
			if tmp.Next == nil {
				break
			}
			tmp = tmp.Next
		}
	}()

	return c
}
