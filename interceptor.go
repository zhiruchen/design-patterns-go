package design

import (
	"log"
	"time"
)

type ReqHandler func(req interface{}) error

type ReqInterceptor func(req interface{}, handler ReqHandler) error

func ChainReqInterceptors(ints ...ReqInterceptor) ReqInterceptor {
	if len(ints) == 0 {
		return func(req interface{}, handler ReqHandler) error {
			return handler(req)
		}
	}

	if len(ints) == 1 {
		return ints[0]
	}

	lastI := len(ints) - 1
	return func(req interface{}, handler ReqHandler) error {
		var chainHandler ReqHandler
		var curI int

		chainHandler = func(req interface{}) error {
			if curI == lastI {
				return handler(req)
			}

			curI++
			err := ints[curI](req, chainHandler)
			curI--
			return err
		}

		return ints[0](req, chainHandler)
	}
}

func LogReqInterceptor(req interface{}, handler ReqHandler) error {
	log.Println("Log Interceptor")
	return handler(req)
}

func AuthReqInterceptor(req interface{}, handler ReqHandler) error {
	log.Println("Auth Interceptor")
	return handler(req)
}

func TimingReqInterceptor(req interface{}, handler ReqHandler) error {
	start := time.Now()
	log.Println("Timing Interceptor")
	defer func() {
		end := time.Now()
		log.Printf("Total Cosuming %fs\n", end.Sub(start).Seconds())
	}()

	return handler(req)
}

func CustomHandler(req interface{}) error {
	log.Println("CustomHandler start")
	time.Sleep(1 * time.Second)
	log.Println(req)
	log.Println("CustomHandler end")
	return nil
}
