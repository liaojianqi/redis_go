/* 链表 */
package main

import (
    "fmt"
)
type Node struct {
    prev *Node
    next *Node
    val interface{}
}
type List struct {
    len int
    head *Node
    tail *Node
}
func (l *List)add_back(v interface{}) {
    if l.len == 0 {
        l.head = new(Node)
        l.head.val = v
        l.tail = l.head
    } else {
        l.tail.next = new(Node)
        l.tail.next.val = v
        l.tail.next.prev = l.tail
        l.tail = l.tail.next
    }
    l.len++
}
func (l *List)show() {
    p := l.head
    for p != nil {
        fmt.Println(p.val)
        p = p.next
    }
}
func main() {
    l := List{0, nil, nil}
    l.add_back("aa")
    l.add_back("bb")
    l.add_back("cc")
    l.add_back("dd")
    l.show()
}
