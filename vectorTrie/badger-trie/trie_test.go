package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	trie := NewTrie()
	trie.Add([]byte("hello"), 1)
	trie.Add([]byte("hello"), 3)
	trie.Add([]byte("hello"), 4)
	trie.Add([]byte("hel"), 20)
	trie.Add([]byte("he"), 20)
	trie.Add([]byte("badger"), 30)

	trie.Add(nil, 10)
	require.Equal(t, map[uint64]struct{}{10: {}}, trie.Get([]byte("A")))

	ids := trie.Get([]byte("hel"))
	require.Equal(t, 2, len(ids))
	require.Equal(t, map[uint64]struct{}{10: {}, 20: {}}, ids)

	ids = trie.Get([]byte("badger"))
	require.Equal(t, 2, len(ids))
	require.Equal(t, map[uint64]struct{}{10: {}, 30: {}}, ids)

	ids = trie.Get([]byte("hello"))
	require.Equal(t, 5, len(ids))
	require.Equal(t, map[uint64]struct{}{10: {}, 1: {}, 3: {}, 4: {}, 20: {}}, ids)

	trie.Add([]byte{}, 11)
	require.Equal(t, map[uint64]struct{}{10: {}, 11: {}}, trie.Get([]byte("A")))

}

func TestTrieDelete(t *testing.T) {
	trie := NewTrie()
	trie.Add([]byte("hello"), 1)
	trie.Add([]byte("hello"), 3)
	trie.Add([]byte("hello"), 4)
	trie.Add(nil, 5)

	trie.Delete([]byte("hello"), 4)

	require.Equal(t, map[uint64]struct{}{5: {}, 1: {}, 3: {}}, trie.Get([]byte("hello")))

	trie.Delete(nil, 5)
	require.Equal(t, map[uint64]struct{}{1: {}, 3: {}}, trie.Get([]byte("hello")))
}