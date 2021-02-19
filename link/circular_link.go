package link

// @Title circular_link
// @Description 循环链表
// 1. 返回给用户的是尾指针，即指向head的指针

// CircularLinkNode 循环列表节点
type CircularLinkNode struct {
	Data interface{}
	Next *CircularLinkNode
}

// CircularLink 循环列表类型
type CircularLink *CircularLinkNode

func createCircularLinkNode(data interface{}) *CircularLinkNode {
	node := new(CircularLinkNode)
	node.Data = data
	node.Next = node
	return node
}

// CircularLinkInit 初始化
func CircularLinkInit(cLink *CircularLink) error {
	node := createCircularLinkNode(nil)
	*cLink = node
	return nil
}

// CircularLinkLength 计算循环链表的长度
func CircularLinkLength(cLink CircularLink) int {
	count := 0

	rear := cLink
	temp := cLink.Next.Next // 第一个元素几点，头结点的下一个节点

	for temp != rear.Next {
		count++
		temp = temp.Next
	}

	return count
}

// CircularLinkInsert 插入节点
func CircularLinkInsert(cLink *CircularLink, i int, data interface{}) error {
	// 新结点是 node要插入3的位置
	// |  1 |  2  |  3  |  4  |  rear |
	//               ^
	//               |
	//               IN
	//
	//  prev  tmp
	// | 1  |  2  |  3  |  4  |  rear |
	//
	//        prev  tmp
	// | 1  |  2  |  3  |  4  |  rear |
	//
	// node.Next = tmp
	// prev.Next = node
	//
	//
	// |  1 |  N  |  2  |  3  |   4   |  rear  |
	//
	//
	node := createCircularLinkNode(data)
	rear := *cLink
	tmp := rear.Next // 头结点
	prev := rear.Next
	step := 0

	for tmp.Next != rear.Next && step < i {
		prev = tmp
		tmp = tmp.Next
		step++
	}

	if tmp.Next == rear.Next { // 如果遍历到最后也就是又指向了头节点
		return ErrOverflow
	}

	node.Next = tmp
	prev.Next = node
	return nil
}

// CircularLinkDelete 删除节点
func CircularLinkDelete(cLink *CircularLink, i int) error {
	// 要删除的node的位置
	// |  1 |  2  |  3  |  4  |  rear |
	//               ^
	//               |
	//              Del
	//
	//  prev  tmp
	// | 1  |  2  |  3  |  4  |  rear |
	//
	//        prev  tmp
	// | 1  |  2  |  3  |  4  |  rear |
	//
	// prev.Next = tmp.Next
	//
	//
	// |  1 |  2  |  4  |  rear  |
	//
	//

	rear := *cLink
	tmp := rear.Next // 头结点
	prev := rear.Next
	step := 0

	for tmp.Next != rear.Next && step < i {
		prev = tmp
		tmp = tmp.Next
		step++
	}

	if tmp.Next == rear.Next { // 如果遍历到最后也就是又指向了头节点
		return ErrOverflow
	}

	prev.Next = tmp.Next

	return nil
}

// GetCircularLinkNode 获取节点
func GetCircularLinkNode(cLink CircularLink, i int, node *CircularLinkNode) error {

	rear := *cLink
	tmp := rear.Next // 头结点
	step := 0

	for tmp.Next != rear.Next && step < i {
		tmp = tmp.Next
		step++
	}

	if tmp.Next == rear.Next { // 如果遍历到最后也就是又指向了头节点
		return ErrOverflow
	}

	*node = *tmp
	return nil
}
