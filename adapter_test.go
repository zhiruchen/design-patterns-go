package design

import "testing"

func TestClientRequestAdapter_Expect(t *testing.T) {
	adapter := &ClientRequestAdapter{whatWeHave: &WhatWeHave{}}
	req := &ClientUserRequest{adapter: adapter}
	req.Request()
}
