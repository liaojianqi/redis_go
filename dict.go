/* 字典 */
/* todo: rehash */
package main

import (
    "fmt"
    "hash/crc32"
)

type Node struct {
    key string
    value interface{}
    next *Node
}
type Dict struct {
    data []*Node
    size uint32
    sizemask uint32
    used uint32
}
func create_dict(size uint32) *Dict {
    d := new(Dict)
    d.size = size
    d.used = 0
    d.sizemask = size - 1
    for i := 0; uint32(i) < size; i++ {
        d.data = append(d.data, nil)
    }
    return d
}
func (d *Dict)hash(key string) uint32 {
    return crc32.ChecksumIEEE([]byte(key)) & d.sizemask
}
func (d *Dict)add_node(key string, value interface{}) {
    node := new(Node)
    node.key = key
    node.value = value
    index := d.hash(key)
    node.next = d.data[index]
    d.data[index] = node
    d.used++
}
func (d *Dict)show() {
    for i := 0; uint32(i) < d.size; i++ {
        node := d.data[i]
        for node != nil {
            fmt.Print(node.key, ":", node.value, " ")
            node = node.next
        }
        fmt.Println("")
    }
}
func main() {
    d := create_dict(4)
    fmt.Println(d.size)
    d.add_node("hello", 100)
    d.add_node("world", 200)
    d.add_node("mal1w1sl", 300)
    d.add_node("more", 400)
    d.add_node("mordsde", 500)
    d.add_node("morefd", 600)
    d.add_node("abcdef", 700)
    d.show()
}
