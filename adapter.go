package design

import "log"

// adapter design pattern: https://dzone.com/articles/design-patterns-uncovered-0

// ClientExpect the interface that client want
type ClientExpect interface {
	Expect()
}

// ClientUserRequestFunc 最终面向用户的接口
type ClientUserRequest struct {
	adapter *ClientRequestAdapter
}

func (req *ClientUserRequest) Request() {
	req.adapter.Expect()
}

// 已经存在的接口
type WhatWeHave struct{}

func (w *WhatWeHave) ExistsExpect() {
	log.Println("System Have ExistsExpect")
}

// 适配器
type ClientRequestAdapter struct {
	whatWeHave *WhatWeHave
}

func (adp *ClientRequestAdapter) Expect() {
	adp.whatWeHave.ExistsExpect()
}
