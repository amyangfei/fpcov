package main

import (
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/amyangfei/fpcov/pkg/hello"
)

func test1(t *testing.T) {
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
	case <-exit:
	}
}

func test2(t *testing.T) {
	var (
		args []string
		exit = make(chan int)
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

	go func() {
		select {
		case <-exit:
			time.Sleep(time.Millisecond * 100)
			os.Exit(0)
		}
	}()

	os.Args = args
	main()
}

func TestRunMain(t *testing.T) {
	// test1(t)
	test2(t)
}
