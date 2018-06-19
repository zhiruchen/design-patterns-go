package design

import "fmt"

type StateContext struct {
	s State
}

func NewStateContext() *StateContext {
	return &StateContext{}
}

func (sc *StateContext) SetState(s State) {
	sc.s = s
}

func (sc *StateContext) Request(val string) {
	sc.s.Handle(sc, val)
}

type State interface {
	Handle(sc *StateContext, val string)
}

type StateWait struct {
}

func (sw *StateWait) Handle(sc *StateContext, val string) {
	fmt.Println(fmt.Sprintf("%s is wait to do", val))
	sc.SetState(&StateWorkInProgress{})
}

type StateWorkInProgress struct {
}

func (swip *StateWorkInProgress) Handle(sc *StateContext, val string) {
	fmt.Println(fmt.Sprintf("%s is work in progress", val))
	sc.SetState(&StateFinished{})
}

type StateFinished struct {
}

func (sf *StateFinished) Handle(sc *StateContext, val string) {
	fmt.Println(val, " is finished ")
	fmt.Println(" end ")
}

func StateChange() {
	sc := NewStateContext()
	sc.SetState(&StateWait{})
	sc.Request("Task")
	sc.Request("Task")
	sc.Request("Task")
	sc.Request("Task")
}
