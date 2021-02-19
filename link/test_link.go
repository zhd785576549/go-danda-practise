package link

import "fmt"

// Student 学生信息
type Student struct {
	Number string
	Name   string
}

// TestLink 测试单项列表
func TestLink() {
	var testLink Link

	Init(&testLink)

	student1 := Student{
		Number: "1111111",
		Name:   "张三",
	}

	student2 := Student{
		Number: "2222222",
		Name:   "李四",
	}

	student3 := Student{
		Number: "3333333",
		Name:   "尾部的人",
	}

	student4 := Student{
		Number: "444444444",
		Name:   "中间插队的人",
	}

	// 添加一个节点
	node1 := new(Node)
	node1.Data = &student1
	node1.Next = nil

	node2 := new(Node)
	node2.Data = &student2
	node2.Next = nil

	if err := AddNode(&testLink, Head, node1, 0); err != nil {
		fmt.Println("Raise a error : ", err)
	}

	if err := AddNode(&testLink, Head, node2, 0); err != nil {
		fmt.Println("Raise a error : ", err)
	}

	node3 := new(Node)
	node3.Data = &student3
	node3.Next = nil
	// 添加到尾部
	if err := AddNode(&testLink, Tail, node3, 0); err != nil {
		fmt.Println("Raise a error : ", err)
	}

	// 添加到中间位置
	node4 := new(Node)
	node4.Data = &student4
	node4.Next = nil
	// 添加到中间
	if err := AddNode(&testLink, Middle, node4, 1); err != nil {
		fmt.Println("Raise a error : ", err)
	}

	// 测试反转
	if err := Reverse(&testLink); err != nil {
		fmt.Println("Raise a error : ", err)
	}

	// 第一种方式，利用协程和channel进行通信异步获取，最坏情况O(n)
	// for node := range LinkGenerator(testLink) {
	// 	stu := interface{}(node.Data).(*Student)
	// 	fmt.Println("--------------Node-------------------")
	// 	fmt.Println("Get student number : ", stu.Number)
	// 	fmt.Println("Get student name : ", stu.Name)
	// 	fmt.Println("--------------Node End-------------------")
	// 	if node.Next == nil {
	// 		break
	// 	}
	// }

	// 第二种方式，首先获取总数，循环获取节点内容，最坏情况 O(n * n)
	if count, err := GetCount(testLink); err == nil {
		for i := 1; i <= count; i++ {
			curNode := new(Node)
			if err = GetNodeByPosition(testLink, i, &curNode); err == nil {
				stu := curNode.Data.(*Student)
				fmt.Println("--------------Node-------------------")
				fmt.Println("Get student number : ", stu.Number)
				fmt.Println("Get student name : ", stu.Name)
				fmt.Println("--------------Node End-------------------")
			}
		}
	}

	// 测试删除
	if err := DeleteNode(&testLink, 3); err != nil {
		fmt.Println("Raise a error : ", err)
	}

	fmt.Println("Delete one node!!!!!")

	if count, err := GetCount(testLink); err == nil {
		for i := 1; i <= count; i++ {
			curNode := new(Node)
			if err = GetNodeByPosition(testLink, i, &curNode); err == nil {
				stu := curNode.Data.(*Student)
				fmt.Println("--------------Node-------------------")
				fmt.Println("Get student number : ", stu.Number)
				fmt.Println("Get student name : ", stu.Name)
				fmt.Println("--------------Node End-------------------")
			}
		}
	}
}

// TestStaticLink 测试静态链表
func TestStaticLink() {
	var testLink StaticLink

	// fmt.Println("Before init addr : ", testLink)

	if err := InitStaticLink(&testLink); err != nil {
		fmt.Println("Raise a error : ", err)
	}

	fmt.Println("After init : ", testLink)

	stu := Student{
		Number: "3443434",
		Name:   "测试1号",
	}

	if err := InsertSLL(testLink, 1, &stu); err != nil {
		fmt.Println("Raise a error : ", err)
	}

	if err := InsertSLL(testLink, 1, &stu); err != nil {
		fmt.Println("Raise a error : ", err)
	}

	fmt.Println("After Insert : ", testLink)
}

// TestCircularLink 测试循环列表
func TestCircularLink() {
	var testLink CircularLink

	CircularLinkInit(&testLink)

	stu1 := Student{
		Number: "1",
		Name:   "测试1号",
	}

	stu2 := Student{
		Number: "2",
		Name:   "测试2号",
	}

	stu3 := Student{
		Number: "3",
		Name:   "测试3号",
	}

	// 测试是否循环，总是打印下一个节点的地址
	// fmt.Println(testLink)
	// fmt.Println(testLink.Next)
	// fmt.Println(testLink.Next)
	// fmt.Println(testLink.Next)
	// fmt.Println(testLink.Next)
	// fmt.Println(testLink.Next)

	CircularLinkInsert(&testLink, 1, &stu1)
	CircularLinkInsert(&testLink, 1, &stu2)
	CircularLinkInsert(&testLink, 1, &stu3)

	count := CircularLinkLength(testLink)
	for i := 0; i < count; i++ {
		var node *CircularLinkNode
		GetCircularLinkNode(testLink, i, node)
		stu := node.Data.(Student)
		fmt.Println("--------------Node-------------------")
		fmt.Println("Get student number : ", stu.Number)
		fmt.Println("Get student name : ", stu.Name)
		fmt.Println("--------------Node End-------------------")
	}
}

// TestStack 测试栈
func TestStack() {
	var s Stack

	// 初始化栈空间
	if err := StackInit(&s); err != nil {
		fmt.Println(err)
	}

}
