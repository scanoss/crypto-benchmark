module example.com/cbench/wrapper/liba

go 1.22

require (
	example.com/cbench/wrapper/libb v1.0.0
)

replace example.com/cbench/wrapper/libb => ../libb
