package main

/*
Automatically set GOMAXPROCS to match Linux container CPU quota.
https://pkg.go.dev/go.uber.org/automaxprocs@v1.6.0
https://github.com/uber-go/automaxprocs
*/
/*
GOMAXPROCS 的默认值是CPU 的核数，这允许Go 程序充分利用机器的每一个CPU，最大程度地提高程序的并发性能。
具体来说：
Go 1.5 之后，GOMAXPROCS 的默认值被设置为CPU 的核数:。
GOMAXPROCS 决定了可以同时运行的用户级Go 代码的系统线程数量上限:。
默认情况下，Go 可以并发执行高达10,000 个线程，但实际并行运行的线程数量受GOMAXPROCS 限制:。
可以通过 runtime.NumCPU() 函数获取当前机器的CPU 核数:。
为什么要这么设置？
Go 程序通常是CPU 密集型的，因此当每个线程可以分配到一个独立的CPU 上执行时，性能最佳。
Go 的调度器会创建goroutine，并将其分配到线程上执行，GOMAXPROCS 限制了可以同时执行的goroutine 数量。
在容器环境中，默认值可能会导致程序使用宿主机的CPU 核数而不是容器分配的CPU 核数，这可能会导致性能问题或CPU 浪费:。
为了解决容器环境的问题，可以使用 uber-go/automaxprocs 或环境变量 GOMAXPROCS 来调整GOMAXPROCS 的值
*/
import _ "go.uber.org/automaxprocs"

func main() {
	// Your application logic here.
}
