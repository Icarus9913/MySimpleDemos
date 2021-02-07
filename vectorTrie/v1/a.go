package main

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

type listHead struct {
	len   int       //长度
	level int       //深度
	root  *trieNode //trie树根节点的引用
}

func (head *listHead) Get(n int) (interface{}, bool) {
	if n < 0 || n >= head.len {
		return nil, false
	}
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
	head.root = setInNode(head.root, n, head.level, value)

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
	return &listHead{0, 0, nil}
}

func main() {

}
