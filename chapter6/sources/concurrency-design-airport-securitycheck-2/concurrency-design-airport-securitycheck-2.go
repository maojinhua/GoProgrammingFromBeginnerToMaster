package main

//  机场例子 并行方案
import "time"

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTmCost = 180
)

func idCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	print("\tgoroutine-", id, ": idCheck ok\n")
	return idCheckTmCost
}

func bodyCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	print("\tgoroutine-", id, ": bodyCheck ok\n")
	return bodyCheckTmCost
}

func xRayCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	print("\tgoroutine-", id, ": xRayCheck ok\n")
	return xRayCheckTmCost
}

func airportSecurityCheck(id int) int {
	print("goroutine-", id, ": airportSecurityCheck ...\n")
	total := 0

	total += idCheck(id)
	total += bodyCheck(id)
	total += xRayCheck(id)

	print("goroutine-", id, ": airportSecurityCheck ok\n")
	return total
}

// 返回接收通道，用于接收各个 goroutine 的耗时数据
func start(id int, f func(int) int, queue <-chan struct{}) <-chan int {
	c := make(chan int)
	go func() {
		total := 0
		for {
			_, ok := <-queue
			if !ok {
				// 通道关闭返回耗时数据
				c <- total
				return
			}
			total += f(id)
		}
	}()
	return c
}

func max(args ...int) int {
	n := 0
	for _, v := range args {
		if v > n {
			n = v
		}
	}
	return n
}

func main() {
	total := 0
	passengers := 30
	queue := make(chan struct{})
	c1 := start(1, airportSecurityCheck, queue)
	c2 := start(2, airportSecurityCheck, queue)
	c3 := start(3, airportSecurityCheck, queue)

	for i := 0; i < passengers; i++ {
		// 向通道发送信号
		queue <- struct{}{}
	}
	// 关闭通道，各个 goroutine 接收到结束信号
	close(queue)

	total = max(<-c1, <-c2, <-c3)
	println("total time cost:", total)
}
