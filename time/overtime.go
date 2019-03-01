package main

import "time"
import "fmt"

func main() {
	// // 在这个例子中，假设我们执行了一个外部调用，2秒之后将结果写入c1
	// c1 := make(chan string, 1)
	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	c1 <- "result 1"
	// }()
	// // 这里使用select来实现超时，`res := <-c1`等待通道结果，
	// // `<- Time.After`则在等待1秒后返回一个值，因为select首先
	// // 执行那些不再阻塞的case，所以这里会执行超时程序，如果
	// // `res := <-c1`超过1秒没有执行的话
	// select {
	// case res := <-c1:
	// 	fmt.Println(res)
	// case <-time.After(time.Second * 1):
	// 	fmt.Println("timeout 1")
	// }

	// // 如果我们将超时时间设为3秒，这个时候`res := <-c2`将在
	// // 超时case之前执行，从而能够输出写入通道c2的值
	// c2 := make(chan string, 1)
	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	c2 <- "result 2"
	// }()
	// select {
	// case res := <-c2:
	// 	fmt.Println(res)
	// case <-time.After(time.Second * 3):
	// 	fmt.Println("timeout 2")
	// }
	// timer()
	// example()
	// test1()
	test2()
	// test3()
}

func timer() {
	// Ticker使用和Timer相似的机制，同样是使用一个通道来发送数据。
	// 这里我们使用range函数来遍历通道数据，这些数据每隔500毫秒被
	// 发送一次，这样我们就可以接收到
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()
	// Ticker和Timer一样可以被停止。一旦Ticker停止后，通道将不再
	// 接收数据，这里我们将在1500毫秒之后停止
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func example() {
	p := fmt.Println
	// 从获取当前时间开始
	now := time.Now()
	p(now)
	// 你可以提供年，月，日等来创建一个时间。当然时间
	// 总是会和地区联系在一起，也就是时区
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	// 你可以获取时间的各个组成部分
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	// 输出当天是周几，Monday-Sunday中的一个 p(then.Weekday())
	// 下面的几个方法判断两个时间的顺序，精确到秒 p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	// Sub方法返回两个时间的间隔(Duration)
	diff := now.Sub(then)
	p(diff)
	// 可以以不同的单位来计算间隔的大小
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	// 你可以使用Add方法来为时间增加一个间隔 // 使用负号表示时间向前推移一个时间间隔 p(then.Add(diff))
	p(then.Add(-diff))
}

// 不同情况下，Timer.Reset()的返回值
func test1() {
	fmt.Println("第1个测试：Reset返回值和什么有关？")
	tm := time.NewTimer(time.Second)
	defer tm.Stop()

	quit := make(chan bool)

	// 退出事件
	go func() {
		time.Sleep(3 * time.Second)
		quit <- true
	}()

	// Timer未超时，看Reset的返回值
	if !tm.Reset(time.Second) {
		fmt.Println("未超时，Reset返回false")
	} else {
		fmt.Println("未超时，Reset返回true")
	}

	// 停止timer
	tm.Stop()
	if !tm.Reset(time.Second) {
		fmt.Println("停止Timer，Reset返回false")
	} else {
		fmt.Println("停止Timer，Reset返回true")
	}

	// Timer超时
	for {
		select {
		case <-quit:
			return

		case <-tm.C:
			if !tm.Reset(time.Second) {
				fmt.Println("超时，Reset返回false")
			} else {
				fmt.Println("超时，Reset返回true")
			}
		}
	}
}

func test2() {
	fmt.Println("\n第2个测试:超时后，不读通道中的事件，可以Reset成功吗？")
	sm2Start := time.Now()
	tm2 := time.NewTimer(time.Second)
	time.Sleep(2 * time.Second)
	fmt.Printf("Reset前通道中事件的数量:%d\n", len(tm2.C))
	if !tm2.Reset(time.Second) {
		fmt.Println("不读通道数据，Reset返回false")
	} else {
		fmt.Println("不读通道数据，Reset返回true")
	}
	fmt.Printf("Reset后通道中事件的数量:%d\n", len(tm2.C))

	select {
	case t := <-tm2.C:
		fmt.Printf("tm2开始的时间: %v\n", sm2Start.Unix())
		fmt.Printf("通道中事件的时间：%v\n", t.Unix())
		if t.Sub(sm2Start) <= time.Second+time.Millisecond {
			fmt.Println("通道中的时间是重新设置sm2前的时间，即第一次超时的时间，所以第二次Reset失败了")
		}
	}

	fmt.Printf("读通道后，其中事件的数量:%d\n", len(tm2.C))
	tm2.Reset(time.Second)
	fmt.Printf("再次Reset后，通道中事件的数量:%d\n", len(tm2.C))
	time.Sleep(2 * time.Second)
	fmt.Printf("超时后通道中事件的数量:%d\n", len(tm2.C))
}

func test3() {
	fmt.Println("\n第3个测试：Reset前清空通道，尽可能通畅")
	smStart := time.Now()
	tm := time.NewTimer(time.Second)
	time.Sleep(2 * time.Second)

	// 停掉定时器再清空
	if !tm.Stop() && len(tm.C) > 0 {
		<-tm.C
	}
	tm.Reset(time.Second)

	// 超时
	t := <-tm.C
	fmt.Printf("tm开始的时间: %v\n", smStart.Unix())
	fmt.Printf("通道中事件的时间：%v\n", t.Unix())
	if t.Sub(smStart) <= time.Second+time.Millisecond {
		fmt.Println("通道中的时间是重新设置sm前的时间，即第一次超时的时间，所以第二次Reset失败了")
	} else {
		fmt.Println("通道中的时间是重新设置sm后的时间，Reset成功了")
	}
}
