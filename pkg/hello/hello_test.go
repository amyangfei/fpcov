package hello

import (
	"testing"

	"github.com/pingcap/failpoint"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	failpoint.Enable("github.com/amyangfei/fpcov/pkg/hello/IfCondInject", "return(true)")
	val := Hello()
	assert.Equal(t, val, success)

	failpoint.Disable("github.com/amyangfei/fpcov/pkg/hello/IfCondInject")
	failpoint.Enable("github.com/amyangfei/fpcov/pkg/hello/IfBodyInject", "return(true)")
	val = Hello()
	assert.Equal(t, val, success2)
}

func TestShake(t *testing.T) {
	Shake()

	defer func() {
		r := recover()
		assert.Equal(t, r, errorOops)
	}()
	failpoint.Enable("github.com/amyangfei/fpcov/pkg/hello/PanicInject", "return(true)")
	Shake()
}

func TestSubRoutineExit(t *testing.T) {
	// Save current function and restore at the end:
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit

	failpoint.Enable("github.com/amyangfei/fpcov/pkg/hello/RoutineExit", "return(true)")
	SubRoutineExit()

	assert.Equal(t, got, 1)
}

func TestBoundary(t *testing.T) {
	failpoint.Enable("github.com/amyangfei/fpcov/pkg/hello/BoundaryEnable", "return(true)")
	Boundary()
}
