package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//无缓冲区的chan在单个协程中会造成死锁
	//c := make(chan int)
	//c <- 1                              // 此时只存在main一个主协程，阻塞在这一步，print没机会执行，造成死锁
	//println(<- c)
	//
	////多个协程成功输出
	//go func() {c <- 1}()              // 此时c会被阻塞 直到值被取走前都不可在塞入新值
	//go func() {println(<-c)}()

	// 带缓存的chan在单个协程中不会阻塞,超过缓冲区大小阻塞
	//c := make(chan int, 1)
	//c <- 1  //在这里阻塞
	//c <- 2
	//println(<- c)

	// chan中没有数据继续读取造成死锁
	//c := make(chan int, 1)
	//c <- 1
	//println(<- c)
	//println(<- c)  //在这里阻塞

	// chan中没有数据继续读取造成死锁,如果关闭后再读取会读到0值，不会死锁
	//c := make(chan int, 1)
	//c <- 1
	//println(<- c)
	//close(c)
	//println(<- c)

	// nil测试
	//c := make(chan error, 1)
	//c <- nil
	//println((<- c).Error())
	//close(c)
	//println(<- c)

	// range读取死锁，rang读取是阻塞式读取
	//chs := make(chan string, 2)
	//chs <- "first"
	//chs <- "second"
	////close(chs)      //没有close会死锁
	//for ch := range chs {
	//	fmt.Println(ch)
	//}
	//
	////select chan
	//MySelect(make(chan int), make(chan int))

	// 接收动作发生在发送动作之前
	//out := make(chan int)
	//out <- 2
	//go f1(out)
	ProductAndConsumer()
}

// select可以同时轮询多个chan，不会造成阻塞
func MySelect(ch chan int, ch2 chan int) {
	for {
		select {
		case a := <-ch:
			fmt.Println(a)
		case b := <-ch2:
			fmt.Println(b)
		}
	}
}

func f1(in chan int) {
	time.Sleep(time.Second)
	fmt.Println(<-in)
}

/**生产者消费者的例子*/
func ProductAndConsumer() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	//带有缓冲区的通道
	cint := make(chan int, 10)
	go func() {
		//product  ，循环往通道中写入一个元素
		for i := 0; i < 100; i++ {
			cint <- i
		}
		//关闭通道
		close(cint)
	}()
	go func() {
		defer wg.Done()
		//consumer   遍历通道消费元素并打印
		for temp := range cint {
			fmt.Println(temp)
			//len函数可以查看当前通道元素个数
			fmt.Println("当前通道元素个数", len(cint))
		}
	}()
	wg.Wait()
}
