package main

import (
	"fmt"
)

/**
golang实现单链表
单链表是一种链式存取的数据结构，用一组地址任意的存储单元存放线性表中的数据元素。
链表中的数据以节点来表示，每个节点的构成：元素(数据元素的映像)+ 指针(指向后继元素存储位置)。元素是存储数据的存储单元，指针是连接每个节点的地址数据。
节点结构：
|--data--|--next--|
data域存放节点值的数据域，next域存放节点的直接后继地址(位置)的指针域(链域)
头指针head和终端节点
单链表中每个几点的存储地址是存放在其前趋节点next域中。而开始节点无前趋，所以应该设置头指针head指向开始节点。链表由头指针唯一确定，单链表可以用头指针的名字来命名。
终端节点无后继，所以终端节点的指针域为空，即NULL。
*/

type Node struct {
	// 值
	Data interface{}
	// 后继节点指针
	Next *Node
}

// 链表是否为空
func IsEmpty(node *Node) bool {
	return node == nil
}

// 是否是最后一个节点
func IsLast(node *Node) bool {
	return node.Next == nil
}

// 查找指定节点的前一个节点
func FindPrevious(data interface{}, node *Node) *Node {
	tmp := node
	for tmp.Next != nil && tmp.Next.Data != data {
		tmp = tmp.Next
	}
	return tmp
}

// 查找某个值在哪个节点
func Find(data interface{}, node *Node) *Node {
	tmp := node
	for tmp.Data != data {
		tmp = tmp.Next
	}
	return tmp
}

// 插入节点:在指定节点插入节点
/**
position:指定的节点位置
*/
func Insert(data interface{}, position *Node) {
	// 新建一个节点
	tmpCell := new(Node)
	if tmpCell == nil {
		fmt.Println("err: out of space")
	}
	// 给新建的节点的值域赋值为传入的data值
	tmpCell.Data = data
	// 新建的节点的next指针指向的是指定节点position的next
	tmpCell.Next = position.Next
	// 指定节点position的后继节点变成了新建的节点
	position.Next = tmpCell
}

// 删除节点
func Delete(data interface{}, node *Node) {
	preview := FindPrevious(data, node)
	tmp := Find(data, node)
	preview.Next = tmp.Next
	tmp = nil
}

// 删除链表
func DeleteList(node **Node) {
	p := node
	for *p != nil {
		tmp := *p
		*p = nil
		*p = tmp.Next
	}
}

//打印列表
func PrintList(list *Node) {
	p := list
	for p != nil {
		fmt.Printf("%d-%v-%p ", p.Data, p, p.Next)
		p = p.Next
	}
	fmt.Println()
}

func main() {
	headNode := &Node{
		Data: 0,
		Next: nil,
	}
	list := headNode
	Insert(1, headNode)
	PrintList(list)
	fmt.Println(IsEmpty(list))
	fmt.Println(IsLast(headNode))
	p := Find(0, list)
	Insert(2, p)
	PrintList(list)
	Delete(1, list)
	PrintList(list)
	DeleteList(&list)
	PrintList(list)
}
