module github.com/tidwall/btree-bench

go 1.26

require (
	github.com/google/btree v1.1.3
	github.com/tidwall/btree v1.8.2
)

require github.com/tidwall/btree/v2 v2.0.0-00010101000000-000000000000 // indirect

replace github.com/tidwall/btree/v2 => ../btree
