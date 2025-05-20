package main

/*
GO 错误处理原则：
1. err != nil 表示有错误
2. 返回错误给上层函数时，需考虑是否需要增加信息以便定位问题，同时考虑是否需要包装错误。
   fmt.Errorf()包装错误使用%w. 不包装使用%v。包装错误意味着错误属于API里面的一部分
   (api提供者需是否合理，是否需要暴露具体实现, 例如，底层用文件实现，后面改用db实现，细节暴露可能不合理)。
3. 包装的错误考虑使用errors.Is()或errors.As()进行错误匹配
*/

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
	main2()
}

func main2() {
	if _, err := os.Open("non-existing"); err != nil {

		err = fmt.Errorf("my wrap err: %w", err)
		/*is preferable to
		if err == fs.ErrExist
		because the former will succeed if err wraps io/fs.ErrExist.
		*/
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}

}

func main3() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

}
