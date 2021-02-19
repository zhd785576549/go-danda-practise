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

// SLLLength 获取长度
func SLLLength(l StaticLink) int {

	count := 0

	// 初始化时如果是空列表，那么最后一个元素的cursor是0
	if l[MaxLinkLength-1].cursor == 0 {
		return 0
	}

	// 获取首个元素的游标
	i := l[MaxLinkLength-1].cursor
	for i != 0 { // 这个跟链表的一路next一样
		i = l[i].cursor
		count++
	}

	return count
}

// InsertSLL 插入静态链表指定位置
func InsertSLL(l StaticLink, i int, data interface{}) error {

	// 找到链表中游标为i的元素
	k := MaxLinkLength - 1
	if i < 1 || i > SLLLength(l)+1 {
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
		// 如果他的next的数据是个空，那么被判定为是最后一个元素，游标被设置为0
		if l[l[k].cursor].data == nil {
			l[j].cursor = 0
		} else {
			l[j].cursor = l[k].cursor
		}
		// 前一个的Next指向，要插入的节点的位置prev.Next = node
		l[k].cursor = j

		// 设置最后一个元素的游标
		l[MaxLinkLength-1].cursor = 1
		return nil
	}

	return ErrUnkownPosition
}

func freeSLL(l StaticLink, i int) error {
	l[i].cursor = l[0].cursor
	l[0].cursor = i
	return nil
}

// DeleteSLL 删除静态链表
func DeleteSLL(l StaticLink, i int) error {
	// | cursor |  6   |  2  |  3  |  4  |  5  |  0  | ... |           1         |
	// | data   |  xxx | xxx | xxx | xxx | xxx | xxx | ... |          nil        |
	// | index  |  0   |  1  |  2  |  3  |  4  |  5  | ... | MaxLinkLength - 1   |
	//
	// 要删除元素是3，倒找cursor是3
	//
	// ------------------------找到
	//                    k     j
	// | cursor |  6   |  2  |  3  |  4  |  5  |  0  | ... |           1         |
	// | data   |  xxx | xxx | xxx | xxx | xxx | xxx | ... |          nil        |
	// | index  |  0   |  1  |  2  |  3  |  4  |  5  | ... | MaxLinkLength - 1   |
	//
	// -----------------变游标---------
	//                    |           |
	// | cursor |  6   |  3  |  3  |  4  |  5  |  0  | ... |           1         |
	// | data   |  xxx | xxx | xxx | xxx | xxx | xxx | ... |          nil        |
	// | index  |  0   |  1  |  2  |  3  |  4  |  5  | ... | MaxLinkLength - 1   |
	//
	// 删除节点
	//
	//-------------赋值给--------
	//             |<-----------|
	// | cursor |  6   |  3  |  6  |  4  |  5  |  0  | ... |           1         |
	// | data   |  xxx | xxx | nil | xxx | xxx | xxx | ... |          nil        |
	// | index  |  0   |  1  |  2  |  3  |  4  |  5  | ... | MaxLinkLength - 1   |
	//
	// -----变换空闲的游标
	//             |
	// | cursor |  2   |  3  |  6  |  4  |  5  |  0  | ... |           1         |
	// | data   |  xxx | xxx | nil | xxx | xxx | xxx | ... |          nil        |
	// | index  |  0   |  1  |  2  |  3  |  4  |  5  | ... | MaxLinkLength - 1   |
	//
	k := MaxLinkLength - 1
	j := 0

	for j = i; j < i-1; j++ {
		k = l[k].cursor
	}

	j = l[k].cursor
	l[k].cursor = l[j].cursor

	if k == MaxLinkLength-1 { // 如果是第一个元素
		l[MaxLinkLength-1].cursor = l[k].cursor
	}

	freeSLL(l, j)
	return nil
}
