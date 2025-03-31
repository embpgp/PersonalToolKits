package chantest

import "fmt"

func send(ch chan int) {
	ch <- 123
}

func recv(ch chan int) {
	v, ok := <-ch
	if !ok {
		fmt.Printf("ok is false")
		return
	}
	fmt.Printf("%v\n", v)
}

func closeSelf(ch chan int) {
	close(ch)
}
