package main

import (
    "fmt"
)

type Node struct {
    a int
}
func get() Node {
    node := Node{10}
    fmt.Println(&node)
    return node
}
func set(node Node) {
    node.a = 100
}
type Dog struct {
    a interface{}
}
func main() {
    v := []int{1,2,3,4}
    d := Dog{v}
    // c := []int(d.a)
    d.a = d.a.([]int)
    fmt.Println(d.a)
    // c[0] = 100
    // fmt.Println(c)
    fmt.Println(d.a)
}
