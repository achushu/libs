package out

// PriorityLevel indicates the severity level of a logged message
type PriorityLevel int

const (
	// PriorityDebug indicates the event is for debugging purposes
	PriorityDebug PriorityLevel = iota
	// PriorityInfo indicates the event is informational in nature
	PriorityInfo
	// PriorityError indicates an error has occurred
	PriorityError
)

// PriorityLevelMap defines string names for PriorityLevels
var PriorityLevelMap = map[string]PriorityLevel{
	"DEBUG": PriorityDebug,
	"INFO":  PriorityInfo,
	"ERROR": PriorityError,
}

func (p PriorityLevel) String() string {
	switch p {
	case PriorityDebug:
		return "DEBUG"
	case PriorityInfo:
		return "INFO"
	case PriorityError:
		return "ERROR"
	}
	return ""
}
