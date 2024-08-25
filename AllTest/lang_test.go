package AllTest

import (
	"testing"
	"time"
)

func TestRoutine(t *testing.T) {
	//go func() {
	//	Routine2()
	//}()

	time.Sleep(time.Second * 10)
}

func TestRoutine2(t *testing.T) {
	go Routine3()
}

func Routine3() {
	for {
		println("hello3")
		time.Sleep(time.Second)
	}
}
