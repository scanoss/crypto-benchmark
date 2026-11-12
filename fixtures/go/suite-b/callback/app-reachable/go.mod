module example.com/cbench/callback/app

go 1.22

require (
	example.com/cbench/callback/liba v1.0.0
	example.com/cbench/callback/libb v1.0.0 // indirect
)

replace example.com/cbench/callback/liba => ../liba

replace example.com/cbench/callback/libb => ../libb
