module example.com/cbench/unused-import/liba

go 1.22

require (
	example.com/cbench/unused-import/libb v1.0.0
)

replace example.com/cbench/unused-import/libb => ../libb
