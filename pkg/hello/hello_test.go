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
	defer func() {
		r := recover()
		assert.Equal(t, r, errorOops)
	}()
	failpoint.Enable("github.com/amyangfei/fpcov/pkg/hello/PanicInject", "return(true)")
	Shake()
}
