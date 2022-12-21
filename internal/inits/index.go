package inits

import (
	"sort"

	"github.com/dezhishen/photo-album-metadata/pkg/config"
)

type initHandler func(cfg *config.Config) error

type handlerWithPriority struct {
	handler  initHandler
	priority uint8
}
type handlerSlice []handlerWithPriority

func (a handlerSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a handlerSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a handlerSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].priority < a[i].priority
}

var allInitHandler handlerSlice

// 优先级
func registerWithPriority(handler initHandler, priority uint8) {
	allInitHandler = append(allInitHandler, handlerWithPriority{
		handler:  handler,
		priority: priority,
	})
	sort.Sort(sort.Reverse(handlerSlice(allInitHandler)))
}

func register(handler initHandler) {
	registerWithPriority(handler, 100)
}

func GetAll() []initHandler {
	var result []initHandler
	for _, e := range allInitHandler {
		result = append(result, e.handler)
	}
	return result
}
