package ghosts

import "sync"

var currentId uint64
var mutex sync.Mutex

func generateGhostId() uint64 {
	mutex.Lock()
	defer mutex.Unlock()

	currentId += 1
	return currentId
}
