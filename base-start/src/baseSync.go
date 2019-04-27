package main

import (
	"fmt"
	"sync"
	"time"
)

/*
同步
*/

var (
	counter = 0
	/*锁*/
	lock sync.Mutex
	/*读写锁*/
	rwLock sync.RWMutex
)

func main() {

	for i := 0; i < 2; i++ {
		go incr()
	}
	time.Sleep(time.Microsecond * 10)
}
func incr() {
	//lock.Lock()
	//rwLock.Lock()
	counter++
	//rwLock.Unlock()
	//defer lock.Unlock()
	fmt.Println(counter)
}
