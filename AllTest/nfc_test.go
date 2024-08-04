package AllTest

import (
	"fmt"
	"github.com/clausecker/nfc/v2"
	"testing"
)

func TestNFC(t *testing.T) {
	//fmt.Println(nfc.Version())
	devices, err := nfc.ListDevices()
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, d := range devices {
		fmt.Printf("found device %s", d)
	}
}
