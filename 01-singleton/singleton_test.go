package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestSingleon(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			p := GetInstance()
			fmt.Printf("p: %v\n", p)
			wg.Done()
		}()
	}
	wg.Wait()
}
