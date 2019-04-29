package hello

import (
	"fmt"

	"github.com/pingcap/failpoint"
)

const (
	success   = 200
	success2  = 201
	errorOops = "ooops..."
)

func setCode(i *int, val int) {
	fmt.Printf("set i from %d to %d!\n", *i, val)
	*i = val
}

func Hello() int {
	var i int
	if func(v *int) {
		failpoint.Inject("IfCondInject", func() {
			fmt.Println("set code inject in if condition")
			setCode(&i, success)
		})
	}(&i); i != success {
		failpoint.Inject("IfCondInject", func() {
			fmt.Println("set code inject in if body")
		})
		setCode(&i, success2)
	}
	return i
}

func Shake() {
	failpoint.Inject("PanicInject", func() {
		panic(errorOops)
	})
	fmt.Println("everything goes well")
}
