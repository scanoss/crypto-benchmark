module example.com/cbench/dispatch/app

go 1.22

require (
	example.com/cbench/dispatch/liba v1.0.0
	example.com/cbench/dispatch/libb v1.0.0 // indirect
)

replace example.com/cbench/dispatch/liba => ../liba

replace example.com/cbench/dispatch/libb => ../libb
