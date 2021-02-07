package main

import "testing"

func BenchmarkListHead_PushBack(b *testing.B) {
	list := New()
	var v interface{}
	for i :=0; i<b.N;i++{
		list.PushBack(v)
	}
}
