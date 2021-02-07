package v3

/*
	NODE_SIZE 是 List 内部节点的宽度，这里我们选用了通用的25，也就是 32 作为 Trie 树节点的宽度。
	这意味着每个 Trie 树节点将会最多有 32 个子节点。
*/

const (
	SHIFT     = 5
	NODE_SIZE = (1 << SHIFT)	//0001 -> 0010 0000= 2^5=32
	MASK      = NODE_SIZE - 1	//31 -> 0001 1111
)