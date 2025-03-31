package chantest

import "testing"

func Test_close(t *testing.T) {
	//var ch chan int
	ch := make(chan int)
	//go send(ch)
	//recv(ch)
	//closeSelf(ch)
	closeSelf(ch)

}
