package AllTest

import (
	"github.com/schollz/progressbar/v3"
	"testing"
	"time"
)

func TestProcessBar(t *testing.T) {
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}

}
