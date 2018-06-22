package design

import (
	"log"
	"testing"
)

func TestChainReqInterceptors(t *testing.T) {
	chain := ChainReqInterceptors(
		LogReqInterceptor,
		AuthReqInterceptor,
		TimingReqInterceptor,
	)

	req := "My REQ"
	err := chain(req, CustomHandler)
	log.Println(err)
}
