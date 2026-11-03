package main

import (
	"github.com/richardartoul/bbtree-go"
	"github.com/tidwall/btree-bench/tools/support"
)

func bench0[T support.KeyType]() {
	var b *bbtree.Map[T, struct{}]
	support.Bench(
		func() { b = new(bbtree.Map[T, struct{}]) },
		func(key T) { support.AssertN2(b.Set(key, struct{}{})) },
		func(key T) { support.Assert2(b.Get(key)) },
	)
}

func main() {
	bench0[int64]()
	bench0[uint64]()
	bench0[string]()
}
