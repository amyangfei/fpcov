package hello

import (
	"testing"

	"github.com/pingcap/failpoint"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	failpoint.Enable("IfCondInject", "return(true)")
	Hello()
	assert.True(t, true)
}

func TestShake(t *testing.T) {
	defer func() {
		r := recover()
		assert.Equal(t, r, errorOops)
	}()
	failpoint.Enable("PanicInject", "return(true)")
	Shake()
}
