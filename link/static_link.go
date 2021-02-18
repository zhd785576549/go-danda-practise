package link

// @Title static_link
// @Description 静态链表联系

// StaticNode 静态链表元素
type StaticNode struct {
	cursor int
	data   interface{}
}

// MaxLinkLength 链表最大长度
var MaxLinkLength int = 100

// StaticLink 静态链表
type StaticLink []StaticNode

// InitStaticLink 初始化静态列表
func InitStaticLink(l *StaticLink) error {
	sLink := make([]StaticNode, MaxLinkLength)
	for i := 0; i < MaxLinkLength-1; i++ {
		sLink[i].cursor = i + 1
	}

	sLink[MaxLinkLength-1].cursor = 0
	*l = sLink
	return nil
}

// mallocSLL 分配一个空间
func mallocSLL(l StaticLink) (int, error) {
	i := l[0].cursor // 第一个元素存放的是空闲游标

	if l[0].cursor != 0 {
		l[0].cursor = l[i].cursor // 分配完了，需要更新备用游标
	}

	return i, nil
}

// InsertSLL 插入静态链表指定位置
func InsertSLL(l StaticLink, i int, data interface{}) error {

	// 找到链表中游标为i的元素
	k := MaxLinkLength - 1
	if i < 1 || i > len(l)+1 {
		return ErrOverflow
	}

	j, err := mallocSLL(l)
	if err != nil {
		return err
	}

	if j != 0 {
		for h := 1; h <= i-1; h++ {
			k = l[k].cursor // 找出前面排序的元素的位置cursor的值，相当于链表的Next找下去
		}
		l[j].data = data
		// 要插入的节点的下一个，改成它前一个的游标的下一个，相当于链表的node.Next = prev.Next
		l[j].cursor = l[k].cursor
		// 前一个的Next指向，要插入的节点的位置prev.Next = node
		l[k].cursor = j
		return nil
	}

	return ErrUnkownPosition
}
