package log

const (
	// event log
	NONE EventLogType = iota
	PASS              // severity : 0
	FAIL              // severity : 0
	// error log
	__ERROR_SEVERITY_MINIMUM__
	DEBUG  EventLogType = iota - 1 // severity : 1
	INFO                           // severity : 2
	NOTICE                         // severity : 3
	WARN                           // severity : 4
	ERR                            // severity : 5
	CRIT                           // severity : 6
	ALERT                          // severity : 7
	EMERG                          // severity : 8
)

var (
	eventLogTypeNames = []string{
		PASS:   "PASS",
		FAIL:   "FAIL",
		DEBUG:  "DEBUG",
		INFO:   "INFO",
		NOTICE: "NOTICE",
		WARN:   "WARN",
		ERR:    "ERR",
		CRIT:   "CRIT",
		ALERT:  "ALERT",
		EMERG:  "EMERG",
	}

	eventLogTypeNameMappingTable = map[string]EventLogType{
		"PASS":   PASS,
		"FAIL":   FAIL,
		"DEBUG":  DEBUG,
		"INFO":   INFO,
		"NOTICE": NOTICE,
		"WARN":   WARN,
		"ERR":    ERR,
		"CRIT":   CRIT,
		"ALERT":  ALERT,
		"EMERG":  EMERG,
	}
)

// EventLogType
type EventLogType int

func (t EventLogType) Severity() int {
	v := t
	if v >= __ERROR_SEVERITY_MINIMUM__ {
		return int(v-__ERROR_SEVERITY_MINIMUM__) + 1
	}
	return 0
}

func (t EventLogType) Name() string {
	return eventLogTypeNames[t]
}

func (t EventLogType) String() string {
	return t.Name()
}

func ParseEventLogTypeName(name string) EventLogType {
	v, ok := eventLogTypeNameMappingTable[name]
	if ok {
		return v
	}
	return NONE
}
