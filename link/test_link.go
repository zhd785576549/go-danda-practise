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
}
