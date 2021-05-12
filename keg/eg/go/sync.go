package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// wait one background goroutine finish
func syncWithMutex() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("hello world!")
		}
		mu.Unlock()
	}()

	mu.Lock()
}

// wait 10 background goroutine finish
func syncWithChannel() {
	done := make(chan int, 10) // 带 10 个缓存

	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("hello world!")
			done <- 1
		}()
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

// wait 10 background goroutine finish with 'wait group'
func syncWithWaitGroup() {
	var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("你好, 世界")
			wg.Done()
		}()
	}

	// 等待N个后台线程完成
	wg.Wait()
}

// sync with producer & consumer
// 生产者
func producer(out chan<- string, tag string) {
	for i := 0; ; i++ {
		out <- fmt.Sprintf("%v : %v", tag, i)
	}
}

// 消费者
func consumer(in <-chan string) {
	for v := range in {
		fmt.Println(v)
	}
}

func syncWithProducerConsumer() {
	ch := make(chan string, 10) // 成果队列

	go producer(ch, "producer 1")
	go producer(ch, "producer 2")
	go consumer(ch)

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

func main() {
	//syncWithMutex()
	syncWithChannel()
}
