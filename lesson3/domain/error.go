package domain

import "errors"

var (
	// ErrNotFound はアイテムが存在しないエラーを表す。
	ErrNotFound = errors.New("Item not found")
	// ErrConflict はサーバーの状態によって操作できないことを表す。
	ErrConflict = errors.New("Conflict with the current state")
)
