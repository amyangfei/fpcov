package hello

import (
	"testing"

	"github.com/pingcap/failpoint"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	failpoint.Enable("github.com/amyangfei/fpcov/pkg/hello/IfCondInject", "return(true)")
	Hello()
	assert.True(t, true)
}

func TestShake(t *testing.T) {
	defer func() {
		r := recover()
		assert.Equal(t, r, errorOops)
	}()
	failpoint.Enable("github.com/amyangfei/fpcov/pkg/hello/PanicInject", "return(true)")
	Shake()
}
