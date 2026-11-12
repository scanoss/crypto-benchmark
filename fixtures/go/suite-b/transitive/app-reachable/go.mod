module example.com/cbench/transitive/app

go 1.22

require (
	example.com/cbench/transitive/liba v1.0.0
	example.com/cbench/transitive/libb v1.0.0 // indirect
)

replace example.com/cbench/transitive/liba => ../liba

replace example.com/cbench/transitive/libb => ../libb
