package sesredis

import (
	"container/list"
	"sync"
)

//session来自redis 实现
type FromRedis struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做 gc
}


