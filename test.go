/*
判断是否为空的单链表

单链表的长度

从头部添加元素

从尾部添加元素

在指定位置添加元素

删除指定元素

删除指定位置的元素

查看是否包含某个元素

遍历所有元素

*/

package main

import (
	"fmt"
)

type ListNode struct {
	Data interface{} //节点数据
	Next *ListNode   //下一个节点数据
}

type LinkedList struct {
	headnode *ListNode //头节点
}

//链表反转
func (L *LinkedList) Reverse() {
	L.headnode = ReverseList(L.headnode)
}

func ReverseList(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	p := ReverseList(pHead.Next)
	pHead.Next.Next = pHead
	pHead.Next = nil
	return p
}

//判断链表是否为空
func (L *LinkedList) IsEmpty() bool {
	if L.headnode == nil {
		return true
	} else {
		return false
	}
}

func (L *LinkedList) Length() int {
	curnode := L.headnode
	count := 0

	for curnode != nil {
		count++
		curnode = curnode.Next
	}

	return count
}

//从头部添加元素
func (L *LinkedList) Add(data interface{}) {

	newnode := &ListNode{Data: data}

	if L.IsEmpty() {
		L.headnode = newnode
	} else {
		curnode := L.headnode
		newnode.Next = curnode
		L.headnode = newnode
	}
}

//从尾部添加节点
func (L *LinkedList) Append(data interface{}) {

	newnode := &ListNode{Data: data}

	if L.IsEmpty() {
		L.headnode = newnode
	} else {
		curnode := L.headnode
		for curnode.Next != nil {
			curnode = curnode.Next
		}
		curnode.Next = newnode
	}
}

//在指定位置添加元素
func (L *LinkedList) Insert(index int, data interface{}) {
	newnode := &ListNode{Data: data}
	if index <= 0 {
		L.Add(data)
	} else if index >= L.Length()-1 {
		L.Append(data)
	} else {
		curnode := L.headnode
		for i := 0; i <= index-1; i++ {
			if i == index-1 {
				newnode.Next = curnode.Next
				curnode.Next = newnode
				return
			}
			curnode = curnode.Next
		}
		curnode = L.headnode
	}
}

//查看是否包含某个元素
func (L *LinkedList) Contain(data interface{}) bool {
	if !L.IsEmpty() {
		curnode := L.headnode
		for {
			if curnode.Data == data {
				return true
			}
			if curnode.Next != nil {
				curnode = curnode.Next
			} else {
				return false
			}
		}
	} else {
		return false
	}
}

//便利所有节点
func (L *LinkedList) Showlist() {
	if !L.IsEmpty() {
		curnode := L.headnode
		for {
			fmt.Printf("\t%v", curnode.Data)
			if curnode.Next != nil {
				curnode = curnode.Next
			} else {
				break
			}
		}
	}
}

//删除指定元素
func (L *LinkedList) Remove(data interface{}) {
	if !L.IsEmpty() {
		pre := L.headnode
		if pre.Data == data {
			L.headnode = pre.Next
		} else {
			for pre.Next != nil {
				if pre.Next.Data == data {
					pre.Next = pre.Next.Next
				} else {
					pre = pre.Next
				}
			}
		}
	}
}

//删除指定位置的元素
func (L *LinkedList) RemoveAtIndex(index int) {

	if index < 0 || index >= L.Length() || L.IsEmpty() {
		return
	}

	pre := L.headnode
	if index == 0 {
		L.headnode = pre.Next
	} else {
		count := 1
		for count != index && pre.Next != nil {
			count++
			pre = pre.Next
		}

		pre.Next = pre.Next.Next
	}
}

func main() {

	list := LinkedList{}

	for i := 1; i <= 5; i++ {
		list.Add(i)
	}

	list.Showlist()

	list.Reverse()
	list.Showlist()

	fmt.Printf("链表长度：%v\n", list.Length())

	list.Insert(11, 10)
	list.Showlist()

	fmt.Printf("是否包含10：%v\n", list.Contain(11))

	list.Remove(1)
	list.Showlist()

}
