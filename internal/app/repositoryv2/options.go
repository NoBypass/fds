package repositoryv2

import "time"

const (
	All          = "*"
	None         = "NONE"
	ReturnDiff   = "DIFF"
	ReturnAfter  = "AFTER"
	ReturnBefore = "BEFORE"
)

type RelateOptions[Edge any] interface {
	contentOption[Edge]
	parallelOption
	timeoutOption
	returnOption
	setOption
}

type DeleteOptions interface {
	timeoutOption
	returnOption
	whereOption
	onlyOption
}

type SelectOptions interface {
	OrderBy(options ...OrderByOptions)
	Fields(fields ...string)
	Range(from, to string)
	Omit(fields ...string)
	Fetch(field string)
	Start(start int64)
	Limit(limit int64)
	parallelOption
	timeoutOption
	whereOption
	onlyOption
	idOption
}

type UpdateOptions[Record any] interface {
	Merge(content *Record)
	contentOption[Record]
	parallelOption
	timeoutOption
	returnOption
	whereOption
	onlyOption
	setOption
	idOption
}

type OrderByOptions interface {
	Desc(field string)
	Asc(field string)
}

type idOption interface {
	ID(ID string)
}

type setOption interface {
	Set(field string, value any)
}

type contentOption[Record any] interface {
	Content(content *Record)
}

type onlyOption interface {
	Only()
}

type returnOption interface {
	Return(fields ...string)
}

type whereOption interface {
	Where(condition string)
}

type timeoutOption interface {
	Timeout(timeout time.Duration)
}

type parallelOption interface {
	Parallel()
}
