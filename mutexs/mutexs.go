package main

import (
	"fmt"
	"sync"
)

type atomic struct {
	value int
	lock  sync.Mutex
}

func (a *atomic) increment() {
	// 锁操作，并在方法结束前释放锁
	// a.lock.Lock()
	// defer a.lock.Unlock()

	a.value++
}

func (a *atomic) get() int {
	// a.lock.Lock()
	// defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomic
	a.increment()

	// 此时若 increment 和 get 方法中未对操作进行 mutex 的控制，可能出现冲突
	// 可使用 go run -race mutexs.go 的命令查看详细的冲突
	go func() {
		a.increment()
	}()

	fmt.Println(a.get())
}
