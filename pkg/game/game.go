package game

import "container/list"

type Board struct {
	pieces list.List
}

type Chess interface {
	is_checkmate() int
	is_stalemate() bool
}