package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type SafeStruct struct {
	mu sync.Mutex
	v  map[string]int
}

func DoMutex() {

	c := SafeStruct{
		v: make(map[string]int),
	}

	for i := 0; i < 15; i++ {
		go c.increment("inc")
	}
	time.Sleep(time.Second)
	fmt.Println("Used mutext and result is :", c.get("inc"))

}

func (s *SafeStruct) increment(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.v[key]++
}

func (s *SafeStruct) get(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.v[key]

}
