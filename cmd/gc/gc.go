package main

import (
	"fmt"
	"runtime"
)

// 使用栈分配（不逃逸）
func noEscape() {
	var x int
	x = 42
	fmt.Println(x)
}

// 使用堆分配（逃逸）
func escapesToHeap() *int {
	x := new(int)
	*x = 42
	return x
}

func main() {
	// 调用不逃逸的函数
	noEscape()

	// 调用逃逸到堆的函数，并打印内存地址
	escapedPtr := escapesToHeap()
	fmt.Printf("Escaped value: %d, Address: %p\n", *escapedPtr, escapedPtr)

	// 打印逃逸分析结果
	// 使用 `go build -gcflags "-m"` 可以查看逃逸分析结果
	// 这里为了示例简洁，不直接在代码中展示分析结果
	// 你可以将上述代码保存为文件，然后使用上述命令编译查看
	// 例如：go build -gcflags "-m" main.go

	// 手动触发GC并等待完成，用于演示（实际使用中不建议这样做）
	runtime.GC()
	fmt.Println("GC completed")
}
