package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"code.byted.org/gopkg/pkg/sync/semaphore"
)

type ChunkItem struct {
	Begin int
	End   int
}

func Chunk(length int, group int) []*ChunkItem {
	if length < 0 || group <= 0 {
		panic("Chunk error:length or group is invalid")
	}

	var result []*ChunkItem
	for begin := 0; begin < length; begin += group {
		end := begin + group
		if end > length {
			end = length
		}
		result = append(result, &ChunkItem{Begin: begin, End: end})
	}

	return result
}

const (
	ChunkSize = 2
	SemaphoreSize = 100
)

func main() {
	var s []int64
	for i := int64(1); i <= 200; i++ {
		s = append(s, i)
	}

	sequential(s)
	//batchConcurrent(s)
}

func batchConcurrent(s []int64) {
	var wg sync.WaitGroup
	sem := semaphore.NewWeighted(SemaphoreSize)
	ctx := context.Background()

	printList(s)

	for _, chunk := range Chunk(len(s), ChunkSize) {
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Println("acquire semaphore error")
		}
		wg.Add(1)
		go func(ck *ChunkItem) {
			defer func() { wg.Done(); sem.Release(1) }()
			res := doSomeCal(s[ck.Begin : ck.End])
			if len(res) != ck.End - ck.Begin {
				fmt.Println("error: return length mismatch")
			}
			// 回填处理结果
			for i := ck.Begin; i < ck.End; i++ {
				s[i] = res[i-ck.Begin]
			}
		}(chunk)
	}

	wg.Wait()
	printList(s)
}

func sequential(s []int64) {
	printList(s)

	for _, ck := range Chunk(len(s), ChunkSize) {
		res := doSomeCal(s[ck.Begin : ck.End])
		if len(res) != ck.End - ck.Begin {
			fmt.Println("error: return length mismatch")
		}
		// 回填处理结果
		for i := ck.Begin; i < ck.End; i++ {
			s[i] = res[i-ck.Begin]
		}
	}

	printList(s)
}

func printList(s []int64) {
	for _, e := range s {
		fmt.Printf("%v ", e)
	}
	fmt.Println("")
}

func doSomeCal(s []int64) []int64 {
	var ret []int64
	for _, e := range s {
		ret = append(ret, 2*e)
	}
	time.Sleep(time.Second)
	return ret
}