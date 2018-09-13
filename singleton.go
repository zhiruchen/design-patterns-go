package design

import "sync"

type singleton struct {
}

var instance *singleton
var once sync.Once

func NewSingleton() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
