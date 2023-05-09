package eroutine

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRoutine(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	results, err := Routine(list, func(t int, c chan<- Result[string]) {
		random := rand.Intn(200)
		time.Sleep(time.Duration(random) * time.Millisecond)
		c <- Result[string]{Result: fmt.Sprintf("获取结果: %d, 随机延迟: %d", t, random)}
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range results {
		t.Log(r)
	}
}
