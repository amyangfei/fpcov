package hello

import (
	"fmt"

	"github.com/pingcap/failpoint"
)

const (
	success   = 200
	errorOops = "ooops..."
)

func setCode(i *int) {
	*i = success
}

func Hello() {
	var i int
	if func(v *int) {
		failpoint.Inject("IfCondInjecct", func() {
			fmt.Println("set code inject in if condition")
		})
		setCode(&i)
	}(&i); i == success {
		failpoint.Inject("IfCondInjecct", func() {
			fmt.Println("set code inject in if body")
		})
		fmt.Printf("i = %d success!\n", i)
	}
}

func Shake() {
	failpoint.Inject("PanicInject", func() {
		panic(errorOops)
	})
	fmt.Println("everything goes well")
}
