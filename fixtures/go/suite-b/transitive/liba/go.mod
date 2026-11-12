module example.com/cbench/transitive/liba

go 1.22

require (
	example.com/cbench/transitive/libb v1.0.0
)

replace example.com/cbench/transitive/libb => ../libb
