module example.com/cbench/deadfn/app

go 1.22

require (
	example.com/cbench/deadfn/liba v1.0.0
	example.com/cbench/deadfn/libb v1.0.0 // indirect
)

replace example.com/cbench/deadfn/liba => ../liba

replace example.com/cbench/deadfn/libb => ../libb
