package main

import "fmt"

const (
	SHIFT     = 5
	NODE_SIZE = (1 << SHIFT) //0001 -> 0010 0000=2^5=32
	MASK      = NODE_SIZE - 1
)

type List interface {
	Get(n int) (interface{}, bool)
	Set(n int, value interface{})
	PushBack(value interface{})
	RemoveBack() interface{}
	Len() int
}

type trieNode struct {
	children []interface{}
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make([]interface{}, NODE_SIZE),
	}
}

func (node *trieNode) getChildNode(index int) *trieNode {
	if child := node.children[index]; nil != child {
		return child.(*trieNode)
	} else {
		return nil
	}
}

func (node *trieNode) getChildValue(index int) interface{} {
	return node.children[index]
}

type listHead struct {
	len   int       //长度
	level int       //深度
	root  *trieNode //trie树根节点的引用
}

func (head *listHead) Get(n int) (interface{}, bool) {
	if n<0 || n >head.len{
		return nil,false
	}
	root := head.root
	for lv := head.level - 1;;lv--{
		index := (n >> uint(lv*SHIFT) & MASK)
		if lv <= 0 {
			//Arrived at leaves node, return value
			return root.getChildValue(index),true
		}else {
			//Update root node
			root = root.getChildNode(index)
		}
	}

}

func (head *listHead) Set(n int, value interface{}) {
	panic("implement me")
}

func (head *listHead) PushBack(value interface{}) {
	panic("implement me")
}

func (head *listHead) RemoveBack() interface{} {
	panic("implement me")
}

func (head listHead) Len() int {
	panic("implement me")
}

func New() List {
	return &listHead{0,0,nil}
}

func main() {
	fmt.Println(NODE_SIZE)
}
