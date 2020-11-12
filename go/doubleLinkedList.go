// package main

// import (
// 	"container/list"
// 	"fmt"
// )

// func main() {
// 	l := list.New()
// 	//末尾插入值为1的元素，并返回该元素。
// 	v1 := l.PushBack(1)
// 	//首部插入值为2的元素，并返回该元素
// 	v2 := l.PushFront(2)
// 	//在元素v1前插入3
// 	l.InsertBefore(3, v2)
// 	//在元素v1后插入4
// 	l.InsertAfter(4, v1)

// 	fmt.Printf("len: %v\n", l.Len())
// 	fmt.Printf("first: %#v\n", l.Front())
// 	fmt.Printf("second: %#v\n", l.Front().Next())
// 	// 遍历列表并打印其内容。
// 	for e := l.Front(); e != nil; e = e.Next() {
// 		fmt.Println(e.Value)
// 	}
// }

package main

import "fmt"

/*
从链表左边取出一个节点:首节点指针右移一位，首节点为nil,长度减一；
给链表末尾新增一个节点:以传入的数据new一个节点，
  1，链表为空时，该node作为链表的首节点、尾节点；
  2，不为空时，该节点赋值到旧的末节点的next，该节点的前置指针为旧的末节点，该节点赋值为链表的末节点，该节点的next为nil，链表长度+1；
反转链表：思路较多，此处随便写写，详见https://lan6193.blog.csdn.net/article/details/104660223
*/
type Node struct {
	Pre  *Node
	Next *Node
	Data string
}

type List struct {
	First *Node
	Last  *Node
	Size  int
}

// 删除链表尾节点
func (list *List) RemoveOneNode() {
	node := list.Last.Pre
	node.Next = nil
	list.Last = node
	list.Size -= 1
}

// 删除链表头节点
func (list *List) RemoveOneNodeByFront() {
	willBeFirstNode := list.First.Next
	willBeFirstNode.Pre = nil
	list.First = willBeFirstNode
	list.Size -= 1
}

// 给链表末尾新增一个节点
func (list *List) AddOneNode(value string) {
	newNode := new(Node)
	newNode.Data = value

	// 链表为空时
	if list.Size < 1 {
		list.First = newNode
		list.Last = newNode
	} else {
		// 链表的长度>=1时
		lastNode := list.Last
		lastNode.Next = newNode
		newNode.Pre = lastNode
		newNode.Next = nil
		list.Last = newNode
	}

	list.Size += 1
}

// 添加新的首节点
func (list *List) AddOneNode2Front(value string) {
	newNode := new(Node)
	newNode.Data = value

	// 链表为空时
	if list.Size < 1 {
		list.First = newNode
		list.Last = newNode
	} else {
		// 链表的长度>=1时
		firstNode := list.First
		firstNode.Pre = newNode
		newNode.Next = firstNode
		list.First = newNode
	}

	list.Size += 1
}

// 创建一个空的双链表
func CreateNewAirList() (list *List) {
	return &List{}
}

// 打印(遍历)链表
func (list *List) Print() {
	currentNode := list.First
	fmt.Print(currentNode.Data, "\t")
	for i := 0; i < list.Size; i++ {
		if currentNode.Next == nil {
			return
		}
		currentNode = currentNode.Next
		fmt.Print(currentNode.Data, "\t")
	}

}

// 逆序打印(遍历)链表
func (list *List) PrintOver() {
	currentNode := &Node{}
	for i := 0; i < list.Size; i++ {
		if currentNode.Next == nil {
			currentNode = list.Last
			fmt.Print(currentNode.Data, "\t")
			currentNode = currentNode.Pre
			continue
		}

		fmt.Print(currentNode.Data, "\t")
		currentNode = currentNode.Pre
	}
}

// 反转链表（意义不大，但还是写了...）
func (list *List) ReverseList(ch chan string) {

	go func() {
		currentNode := &Node{}
		for i := 0; i < list.Size; i++ {
			if currentNode.Next == nil {
				currentNode = list.Last
				ch <- currentNode.Data
				currentNode = currentNode.Pre
				continue
			}

			if currentNode.Pre == nil {
				ch <- currentNode.Data
				close(ch)
				return
			}

			ch <- currentNode.Data
			currentNode = currentNode.Pre
		}
	}()

	for value := range ch {
		list.AddOneNode(value)
	}

	for i := 0; i < list.Size; i++ {
		list.RemoveOneNodeByFront()
	}

}

func main() {
	list := CreateNewAirList()
	list.AddOneNode("A") // 添加元素
	list.AddOneNode("B")
	list.AddOneNode("C")
	fmt.Println(list.First.Data, list.First.Next.Data, list.Last.Data) // 打印方式1
	list.Print()                                                       // 打印方式2
	fmt.Println("\n给开头添加一个元素T：")
	list.AddOneNode2Front("T") // 给开头添加一个元素
	list.Print()
	list.RemoveOneNode() // 取下尾节点
	fmt.Println("\n取下尾节点后：")
	list.Print()
	list.RemoveOneNodeByFront() // 取下首节点
	fmt.Println("\n取下首节点后：")
	list.Print()
	list.AddOneNode("C")
	list.AddOneNode("D")
	fmt.Println("\n添加元素C、D并打印：")
	list.Print()

	fmt.Println("\n对链表进行反转：")
	listCh := make(chan string)
	list.ReverseList(listCh)
	list.Print()

}
