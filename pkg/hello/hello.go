package hello

import (
	"fmt"
	"sync"

	"github.com/pingcap/failpoint"

	"github.com/amyangfei/fpcov/pkg/atexit"
)

const (
	success   = 200
	success2  = 201
	errorOops = "ooops..."
)

func init() {
	atexit.Register(func() {})
}

func setCode(i *int, val int) {
	fmt.Printf("set i from %d to %d!\n", *i, val)
	*i = val
}

func Hello() int {
	var i int
	if j := func(v *int) int {
		failpoint.Inject("IfCondInject", func() {
			fmt.Println("set code inject in if condition")
			setCode(&i, success)
		})
		return *v
	}(&i); j != success {
		failpoint.Inject("IfBodyInject", func() {
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

func SubRoutinePanic() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		failpoint.Inject("RoutinePanic", func() {
			fmt.Println("exit injection")
			atexit.Exit(0)
		})
	}()
	wg.Wait()
}
