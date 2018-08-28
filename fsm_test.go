package design

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateMachine(t *testing.T) {
	startTransition := func(args ...interface{}) (string, error) {
		v, ok := args[0].(string)
		if !ok {
			return "error_state", nil
		}

		if v != "Golang" {
			return "error_state", nil
		}

		return "is_state", nil
	}

	isTransition := func(args ...interface{}) (string, error) {
		v, ok := args[1].(string)
		if !ok {
			return "error_state", nil
		}

		if v != "is" {
			return "error_state", nil
		}

		return "concise_state", nil
	}

	conciseTransition := func(args ...interface{}) (string, error) {
		v, ok := args[2].(string)
		if !ok {
			return "error_state", nil
		}

		if v != "concise" && v != strings.ToUpper("concise") {
			return "error_state", nil
		}

		return "end_state", nil
	}

	s := NewStateMachine("start_state")
	s.AddState("start_state", startTransition, false)
	s.AddState("is_state", isTransition, false)
	s.AddState("concise_state", conciseTransition, false)
	s.AddState("end_state", nil, true)
	s.AddState("error_state", nil, true)

	cases := []struct {
		args   []interface{}
		st     string
		result error
	}{
		{
			args:   []interface{}{"Golang", "is", "concise"},
			st:     "END_STATE",
			result: nil,
		},
		{
			args:   []interface{}{"golang", "is", "concise"},
			st:     "ERROR_STATE",
			result: nil,
		},
		{
			args:   []interface{}{"Golang", "is", "CONCISE"},
			st:     "END_STATE",
			result: nil,
		},
		{
			args:   []interface{}{"Golang", "IS", "CONCISE"},
			st:     "ERROR_STATE",
			result: nil,
		},
	}

	ast := assert.New(t)
	for _, c := range cases {
		st, err := s.Run(c.args...)
		ast.Equal(c.st, st, "state not equal")
		ast.Nil(err, "")
	}
}
