package program

import "sync"

// 单例模式
// https://refactoringguru.cn/design-patterns/singleton/go/example

type single struct{}

var singleInstance *single

// 解法一
var lock = &sync.Mutex{}

func GetInstance1() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{}
		}
	}

	return singleInstance
}

// 解法二
var once sync.Once

func GetInstance2() *single {
	once.Do(func() {
		singleInstance = &single{}
	})

	return singleInstance
}
