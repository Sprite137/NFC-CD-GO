package enum

type LogicType int

const (
	ORDER = iota
	RANDOM
	LOOP
)

func (i LogicType) String() string {
	switch i {
	case 0:
		return "ORDER"
	case 1:
		return "RANDOM"
	case 2:
		return "LOOP"
	}
	return "N/A"
}
