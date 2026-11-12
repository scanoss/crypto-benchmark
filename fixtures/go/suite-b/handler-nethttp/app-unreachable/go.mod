module example.com/cbench/handler-nethttp/app

go 1.22

require (
	example.com/cbench/handler-nethttp/liba v1.0.0
	example.com/cbench/handler-nethttp/libb v1.0.0 // indirect
)

replace example.com/cbench/handler-nethttp/liba => ../liba

replace example.com/cbench/handler-nethttp/libb => ../libb
