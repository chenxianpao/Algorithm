package main

import (
	"fmt"
	"log"
	"reflect"
)

/*
栈的基本操作：
- 栈的初始化 InitStack
- 压栈 Push
- 弹栈 Pop
- 栈大小 StackLen
- 栈的容量 StackSize
- 判断是否为空 IsEmpty
- 判断是否已满 IsFull
- 获取栈顶元素的值 Peek
- 清空栈 Clear
- 查询某个值==最近==距离栈顶的距离 Search
- 遍历栈 Traverse

*/

type stack struct {
	size   int64         //容量
	length int64         //栈的长度
	data   []interface{} //存在的元素
}

//初始化栈
func InitStack(size int64) stack {
	t := stack{}
	t.size = size
	t.data = make([]interface{}, size)
	return t
}

//压栈
func (s *stack) Push(r interface{}) (b bool, err error) {
	if s.IsFull() {
		b = false
		err = fmt.Errorf("Stack is Full")
		log.Println("Stack is Full")
		return
	}
	s.data[s.length] = r
	s.length++
	return
}

//弹栈
func (s *stack) Pop() (r interface{}, err error) {
	if s.IsEmpty() {
		err = fmt.Errorf("Stack is Empty")
		log.Println("Stack is Empty")
		return
	}
	s.length--
	r = s.data[s.length]
	return
}

//获取栈顶对象
func (s *stack) Peek() (r interface{}, err error) {
	if s.IsEmpty() {
		err = fmt.Errorf("Stack is Empty")
		log.Printf("Stack is Empty")
		return
	}
	r = s.data[s.length-1]
	return
}

/*
	搜索某一个item在该栈中的位置【位置为离栈顶最近的item与栈顶间距离】
	方法调用返回从堆栈中，对象位于顶部的基于1的位置。
*/
func (s *stack) Search(r interface{}) (post int64, err error) {
	post = int64(0)
	if s.IsEmpty() {
		err = fmt.Errorf("")
		log.Printf("Stack is Empty")
		return
	}

	for k, v := range s.data {
		if reflect.DeepEqual(v, r) {
			post = int64(s.length) - int64(k)
		}
	}
	return
}

//遍历栈
func (s *stack) Traverse(fn func(r interface{}), isTop2Bottom bool) {
	if isTop2Bottom {
		var i int64 = 0
		for ; i < s.length; i++ {
			fn(s.data[i])
		}
	} else {
		for i := s.length - 1; i >= 0; i-- {
			fn(s.data[i])
		}
	}
}

//栈的容量
func (s *stack) StackSize() int64 {
	return s.size
}

//栈的长度
func (s *stack) StackLen() int64 {
	return s.length
}

//判断是否栈为空
func (s *stack) IsEmpty() bool {
	return s.length == 0
}

//判断是否已经满了
func (s *stack) IsFull() bool {
	return s.length == s.size
}

//清空
func (s *stack) Clear() {
	s.length = 0
	s.data = s.data[0:0]
}

// ----------------------
// Stack 是用于存放 int 的 栈
type Stack struct {
	nums []int
}

// NewStack 返回 *kit.Stack
func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

// Push 把 n 放入 栈
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

// Pop 从 s 中取出最后放入 栈 的值
func (s *Stack) Pop() int {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

// Len 返回 s 的长度
func (s *Stack) Len() int {
	return len(s.nums)
}

// IsEmpty 反馈 s 是否为空
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
