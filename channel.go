package main

func main() {
	// 정수형 채널을 생성
	ch := make(chan int)
	go func() {
		ch <- 123
	}()

	var i int
	i = <-ch
	println(i)

}
