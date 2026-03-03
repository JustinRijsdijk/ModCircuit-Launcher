package repositoryStatus

type Status int

const (
	StatusUnknown Status = iota
	StatusOnline
	StatusOffline
	StatusSyncing
)

func (s Status) String() string {
	switch s {
	case StatusOnline:
		return "online"
	case StatusOffline:
		return "offline"
	case StatusSyncing:
		return "syncing"
	case StatusUnknown:
		fallthrough
	default:
		return "unknown"
	}
}

func ParseStatus(value string) Status {
	switch value {
	case "online":
		return StatusOnline
	case "offline":
		return StatusOffline
	case "syncing":
		return StatusSyncing
	default:
		return StatusUnknown
	}
}