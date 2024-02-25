package singleton

import (
	"sync"
	"time"
)

// 两种模式:
// 1.饿汉式模式: 在包加载的时候把结构体给创建了
// type Person struct {
// 	t time.Time
// }

// var p *Person

// func init() {
// 	p = &Person{t: time.Now()}
// }

// func GetInstance() *Person {
// 	return p
// }

// 2.懒汉式模式: 只有结构体用到时候才创建结构体
type Person struct {
	t time.Time
}

var lazyPerson *Person
var once = sync.Once{} // 只执行一次

func GetInstance() *Person {
	// Do 中的函数只有第一次调用时执行 之后的会忽略
	once.Do(func() {
		lazyPerson = &Person{t: time.Now()}
	})
	return lazyPerson
}
