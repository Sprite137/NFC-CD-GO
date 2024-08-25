package util

import (
	"encoding/hex"
	"errors"
	"github.com/clausecker/nfc/v2"
)

func GetNfcUID(target nfc.Target) (*string, error) {

	var UID string
	// Transform the target to a specific tag Type and send the UID to the channel
	switch target.Modulation() {
	case nfc.Modulation{Type: nfc.ISO14443a, BaudRate: nfc.Nbr106}:
		var card = target.(*nfc.ISO14443aTarget)
		var UIDLen = card.UIDLen
		var ID = card.UID
		UID = hex.EncodeToString(ID[:UIDLen])
		break
	case nfc.Modulation{Type: nfc.ISO14443b, BaudRate: nfc.Nbr106}:
		var card = target.(*nfc.ISO14443bTarget)
		var UIDLen = len(card.ApplicationData)
		var ID = card.ApplicationData
		UID = hex.EncodeToString(ID[:UIDLen])
		break
	case nfc.Modulation{Type: nfc.Felica, BaudRate: nfc.Nbr212}:
		var card = target.(*nfc.FelicaTarget)
		var UIDLen = card.Len
		var ID = card.ID
		UID = hex.EncodeToString(ID[:UIDLen])
		break
	case nfc.Modulation{Type: nfc.Felica, BaudRate: nfc.Nbr424}:
		var card = target.(*nfc.FelicaTarget)
		var UIDLen = card.Len
		var ID = card.ID
		UID = hex.EncodeToString(ID[:UIDLen])
		break
	case nfc.Modulation{Type: nfc.Jewel, BaudRate: nfc.Nbr106}:
		var card = target.(*nfc.JewelTarget)
		var ID = card.ID
		var UIDLen = len(ID)
		UID = hex.EncodeToString(ID[:UIDLen])
		break
	case nfc.Modulation{Type: nfc.ISO14443biClass, BaudRate: nfc.Nbr106}:
		var card = target.(*nfc.ISO14443biClassTarget)
		var ID = card.UID
		var UIDLen = len(ID)
		UID = hex.EncodeToString(ID[:UIDLen])
		break
	default:
		return nil, errors.New("unknown modulation")
	}

	return &UID, nil
}
