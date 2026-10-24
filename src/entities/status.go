package entities

type Status uint8

const (
	UP = iota
	DOWN
	DEGRADED
	MAINTENANCE
)

func (s Status) String() string {
	switch s {
	case DOWN:
		return "DOWN"
	case UP:
		return "UP"
	case DEGRADED:
		return "DEGRADED"
	case MAINTENANCE:
		return "MAINTENANCE"
	default:
		return "UNKNOWN"
	}
}
