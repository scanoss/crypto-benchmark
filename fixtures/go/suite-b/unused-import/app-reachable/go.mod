module example.com/cbench/unused-import/app

go 1.22

require (
	example.com/cbench/unused-import/liba v1.0.0
	example.com/cbench/unused-import/libb v1.0.0 // indirect
)

replace example.com/cbench/unused-import/liba => ../liba

replace example.com/cbench/unused-import/libb => ../libb
