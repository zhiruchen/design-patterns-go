package design

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewComplexObject(t *testing.T) {
	obj := NewComplexObject(
		WithOp1("OPTION1"),
		WithOp2(42),
		WithOp3(42.234567897654),
		WithOp4(map[string]interface{}{"1234": 1234}),
		WithOp5("OPTION5"),
	)

	result := &Options{
		op1: "OPTION1",
		op2: 42,
		op3: 42.234567897654,
		op4: map[string]interface{}{"1234": 1234},
		op5: "OPTION5",
	}

	assert.True(t, reflect.DeepEqual(obj.opts, result))
}
