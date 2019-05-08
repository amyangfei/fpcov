package hello

import (
	"fmt"
)

func Goodbye() {
	defer func() {
		fmt.Println("do some cleanup")
	}()

	fmt.Println("before exit halfway")
	OsExit(1)
	fmt.Println("after exit halfway")
}
