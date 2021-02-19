package link

import "errors"

// @Title errno
// @Description 错误代码
// ErrEmptyLink 空连表
var (
	ErrEmptyLink      = errors.New("empty link")
	ErrAlreadyInit    = errors.New("already init")
	ErrOverflow       = errors.New("node add over flow the link")
	ErrUnkownPosition = errors.New("unknown position")
	ErrFull           = errors.New("space already full")
	ErrEmpty          = errors.New("space is empty")
)
