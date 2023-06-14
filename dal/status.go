package dal

const (
	WAIT = 1 + iota
	SCHEDULED
	EXECUTING
	SUCCESS
	FAILED
)
