package hello

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/pingcap/failpoint"
)

const (
	success   = 200
	success2  = 201
	errorOops = "ooops..."
)

var (
	osExit func(int)
)

func init() {
	osExit = os.Exit
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

func SubRoutineExit() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		failpoint.Inject("RoutineExit", func() {
			fmt.Println("exit injection")
			osExit(1)
		})
	}()
	wg.Wait()
}

func Boundary() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 100)
	}()
	go func() {
		fmt.Println("before sleep")
		time.Sleep(time.Second * 300)
		fmt.Println("after sleep")
	}()
	wg.Wait()
}
