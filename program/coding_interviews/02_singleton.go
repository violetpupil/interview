package program

import "sync"

// 单例模式
// https://refactoringguru.cn/design-patterns/singleton/go/example

type single struct{}

var singleInstance *single

var lock = &sync.Mutex{}

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

// golang实现
var once sync.Once

func Golang() *single {
	once.Do(func() {
		singleInstance = &single{}
	})

	return singleInstance
}
