package AllTest

import (
	"testing"
)

const (
	NDEF_RECORD_HEADER_LEN = 3
)

type NDEFRecord struct {
	Header  uint8
	TypeLen uint8
	IDLen   uint8
	Payload []byte
}

func TestTxt2dump(t *testing.T) {

}
