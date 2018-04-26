package resourcebase

import (
	"testing"
	"time"
	"sync"
)

// This test will run 11 concurrent tasks, each task takes 5 seconds.
// We set the size of resourcebase to 10 which means the addition 1 task will have to 
// wait for the 1 of the first 10 tasks to finish.
// We check whether the whole process takes longer than 10 seconds.
func TestResourceBase(t *testing.T) {
	rb := NewResourceBase("test", 10)
	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 11; i++ {
		wg.Add(1)
		go func() {
			Query(rb)
			wg.Done()
		}()
	}
	wg.Wait()
	duration := time.Since(start)
	if duration < time.Second * 10 {
		t.Fatalf("Result does not meet expectation! Duration should be greater than 10 seconds")
	}
}

func Query(r ResourceBase) {
	r.Take()
	defer r.Return()
	time.Sleep(time.Second * 5)
}