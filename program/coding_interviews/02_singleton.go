package program

import "sync"

// 单例模式
// https://refactoringguru.cn/design-patterns/singleton/go/example

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{}
		}
	}

	return singleInstance
}
