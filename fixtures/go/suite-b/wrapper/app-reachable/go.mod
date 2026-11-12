module example.com/cbench/wrapper/app

go 1.22

require (
	example.com/cbench/wrapper/liba v1.0.0
	example.com/cbench/wrapper/libb v1.0.0 // indirect
)

replace example.com/cbench/wrapper/liba => ../liba

replace example.com/cbench/wrapper/libb => ../libb
