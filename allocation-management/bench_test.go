package main

import "testing"

type T struct {
	X [1000]int32 // 4B
}

var global interface{}

func BenchmarkAllocOnHeap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i <= b.N; i++ {
		global = &T{}
	}
}

func BenchMarkAllocOnStack(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		local := T{}
		_ = local
	}
}
