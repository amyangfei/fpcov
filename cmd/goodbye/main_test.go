package main

import (
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/amyangfei/fpcov/pkg/hello"
)

func TestRunMain(t *testing.T) {
	var (
		args   []string
		exit   = make(chan int)
		waitCh = make(chan interface{}, 1)
	)
	for _, arg := range os.Args {
		switch {
		case arg == "DEVEL":
		case strings.HasPrefix(arg, "-test."):
		default:
			args = append(args, arg)
		}
	}

	oldOsExit := hello.OsExit
	defer func() { hello.OsExit = oldOsExit }()
	hello.OsExit = func(code int) {
		log.Printf("[test] os.Exit with code %d", code)
		exit <- code
		// sleep here to prevent following code execution in the caller routine
		time.Sleep(time.Second * 60)
	}

	os.Args = args
	go func() {
		main()
		close(waitCh)
	}()

	select {
	case <-waitCh:
		return
	case <-exit:
		return
	}
}
