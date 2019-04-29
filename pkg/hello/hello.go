package hello

import (
	"fmt"
	"github.com/pingcap/failpoint"
)

func setCode(i *int) {
	*i = 200
}

const success = 200

func Hello() {
	var i int
	if func(v *int) {
		failpoint.Inject("fptest-in-pkg", func() {
			fmt.Println("set code inject in if condition")
		})
		*v = success
	}(&i); i == success {
		failpoint.Inject("fptest-in-pkg", func() {
			fmt.Println("set code inject in if body")
		})
		fmt.Printf("i = %d success!\n", i)
	}
}
