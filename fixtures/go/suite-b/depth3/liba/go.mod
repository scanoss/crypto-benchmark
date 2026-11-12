module example.com/cbench/depth3/liba

go 1.22

require (
	example.com/cbench/depth3/libb v1.0.0
	example.com/cbench/depth3/libc v1.0.0 // indirect
)

replace example.com/cbench/depth3/libb => ../libb

replace example.com/cbench/depth3/libc => ../libc
