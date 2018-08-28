package design

import (
	"fmt"
	"strings"
)

type TransitionsHandler func(args ...interface{}) (string, error)

type stateMachine struct {
	handlers   map[string]TransitionsHandler
	startState string
	endStates  map[string]struct{}
}

func NewStateMachine(startState string) *stateMachine {
	return &stateMachine{
		handlers:   make(map[string]TransitionsHandler),
		endStates:  make(map[string]struct{}, 1),
		startState: strings.ToUpper(startState),
	}
}

func (m *stateMachine) AddState(name string, h TransitionsHandler, endState bool) {
	name = strings.ToUpper(name)
	m.handlers[name] = h
	if endState {
		m.endStates[name] = struct{}{}
	}
}

func (m *stateMachine) Run(args ...interface{}) (string, error) {
	handler, ok := m.handlers[m.startState]
	if !ok {
		return "", fmt.Errorf("no start state")
	}

	if len(m.endStates) == 0 {
		return "", fmt.Errorf("no end states")
	}

	for {
		newState, err := handler(args...)
		if err != nil {
			return "", err
		}

		newState = strings.ToUpper(newState)
		if _, ok := m.endStates[newState]; ok {
			return newState, nil
		}

		handler = m.handlers[newState]
	}

	return "", nil
}
