package main

import "fmt"

const (
	SHIFT     = 5
	NODE_SIZE = (1 << SHIFT)  //0001 -> 0010 0000=2^5=32
	MASK      = NODE_SIZE - 1 //31 -> 0001 1111
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

func newTrieNode() *trieNode { //切片的使用: 1.从数组中抽出来  2.直接切片  3.make初始化
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

/*
	在Get、Set操作之时，我们先检查目标 Index 是否大于等于offset，如果为真，我们就直接在 tail 节点上进行操作。
	否则说明目标元素在 Trie 树当中，我们仍然使用之前的 Trie 树操作。
*/
type listHead struct {
	len    int       //切片长度
	level  int       //深度
	root   *trieNode //trie树根节点的引用
	offset int       //代表的是列表中在tail节点之前的节点当中存储的元素的数量,同时也是tail节点中下标0的元素在整个List当中的Index
	tail   *trieNode //尾巴
}

func (head *listHead) Get(n int) (interface{}, bool) {
	if n < 0 || n >= head.len {
		return nil, false
	}
	if n >= head.offset {
		return head.tail.getChildValue(n - head.offset), true
	}
	//Get elements in the trie
	root := head.root
	for lv := head.level - 1; ; lv-- {
		index := (n >> uint(lv*SHIFT) & MASK)
		if lv <= 0 {
			//Arrived at leaves node, return value
			return root.getChildValue(index), true
		} else {
			//Update root node
			root = root.getChildNode(index)
		}
	}
}

func (head *listHead) Set(n int, value interface{}) {
	if n < 0 || n >= head.len {
		panic("Index out of bound")
	}
	if n >= head.offset {
		head.tail = setTail(head.tail, n-head.offset, value)
	} else {
		head.root = setInNode(head.root, n, head.level, value)
	}
}

func setTail(tail *trieNode, n int, value interface{}) *trieNode {
	if nil == tail {
		tail = newTrieNode()
	}
	tail.children[n] = value
	return tail
}

/*
	index := (n >> uint(level-1)*SHIFT) & MASK 解释:
	根据level值计算当前应该用来查询子元素Si,也就是目标子元素在数据中的位置
	其中SHIFT=5, MASK=(1<<SHIFT)-1=31, MASK的二进制表示从最低位开始向上恰好是5个1
	这样我们就把n每5位一组分为一个Symbol进行查询了
*/
func setInNode(root *trieNode, n int, level int, value interface{}) *trieNode {
	index := (n >> uint(level-1) * SHIFT) & MASK
	if level == 1 {
		root.children[index] = value
	} else {
		child := root.getChildNode(index)
		root.children[index] = setInNode(child, n, level-1, value)
	}
	return root
}

/*
	目前只支持从数组尾部添加和删除元素.
	如果在试图进行PushBack的时候tail中的元素已满,需要将当前的tail节点放入Trie中,维护和更新offset的值,然后新建一个tail出来并把元素插入到新的tail上
	offset是tail当中第一个元素在List中的位置
	我们可以在Trie树中找出offset位置的元素应该存在的位置,那里自然也是tail应该被放置的地方
*/
func (head *listHead) PushBack(value interface{}) {
	//Increase the depth of tree while the capacity is not enough
	if head.len-head.offset < NODE_SIZE {
		//Tail node has free space
		head.tail = setTail(head.tail, head.len-head.offset, value)
	} else {
		//Tail node is full
		n := head.offset
		lv := head.level
		root := head.root

		for lv == 0 || (n>>uint(lv*SHIFT)) > 0 {
			parent := newTrieNode()
			parent.children[0] = root
			root = parent
			lv++
		}
		head.root = putTail(root, head.tail, n, lv)
		head.tail = nil
		head.tail = setTail(head.tail, 0, value)

		head.level = lv
		head.offset += NODE_SIZE
	}
	head.len++
}

func putTail(root *trieNode, tail *trieNode, n int, level int) *trieNode {
	index := (n >> uint(level-1) * SHIFT) & MASK
	if nil == root {
		root = newTrieNode()
	}
	if level == 1 {
		return tail
	} else {
		root.children[index] = putTail(root.getChildNode(index), tail, n, level-1)
	}
	return root
}

func (head *listHead) RemoveBack() interface{} {
	if head.len == 0 {
		panic("Remove from empty list")
	}
	value := head.tail.getChildValue(head.len - head.offset - 1)
	head.tail = setTail(head.tail, head.len-head.offset-1, nil) //clear reference to release memory

	head.len--

	if head.len == 0 {
		head.level = 0
		head.offset = 0
		head.root = nil
		head.tail = nil
	} else {
		if head.len <= head.offset {
			//tail is empty, retrieve new tail from root
			head.root, head.tail = getTail(head.root, head.len-1, head.level)
			head.offset -= NODE_SIZE
		}

		//Reduce the depth of tree if root only have one child
		n := head.offset - 1
		lv := head.level
		root := head.root

		for lv > 1 && (n>>uint(lv-1)*SHIFT) == 0 {
			root = root.getChildNode(0)
			lv--
		}
		head.root = root
		head.level = lv
	}
	return value
}

func (head listHead) Len() int {
	return head.len
}

func New() List {
	return &listHead{len: 0, level: 0, root: nil, offset: 0, tail: nil}
}

func getTail(root *trieNode, n int, level int) (*trieNode, *trieNode) {
	index := (n >> uint(level-1) * SHIFT) & MASK
	if level == 1 {
		return nil, root
	} else {
		child, tail := getTail(root.getChildNode(index), n, level)
		if index == 0 && child == nil {
			//The first element has been removed, which means current node
			//becomes empty, remove current node by returning nil
			return nil, tail
		} else {
			//Current node is not empty
			return root, tail
		}
	}
}

func main() {
	list := New()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	list.PushBack(6)
	list.PushBack(7)

	fmt.Println(list)

}
