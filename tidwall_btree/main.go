package main

import (
	"github.com/tidwall/btree-bench/tools/support"
	"github.com/tidwall/btree/v2"
)

func bench0[T support.KeyType]() {
	var b *btree.Map[T, struct{}]
	support.Bench(
		func() { b = btree.NewMap[T, struct{}]() },
		func(key T) { support.Assert2(b.Insert(key, struct{}{})) },
		func(key T) { support.Assert2(b.Lookup(key)) },
	)
}

func main() {
	bench0[int32]()
	bench0[uint64]()
	bench0[string]()
}
