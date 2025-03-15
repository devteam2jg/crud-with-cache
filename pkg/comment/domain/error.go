package domain

import (
	"errors"
)

var (
	ErrCommentNotFound = errors.New("comment not found")
	ErrCommentOwner    = errors.New("comment owner not match")

	ErrFeedNotFound = errors.New("feed not found")
	ErrFeedOwner    = errors.New("feed owner not match")
)

func Validate(err error) bool {
	if errors.Is(err, ErrCommentNotFound) {
		return true
	}
	if errors.Is(err, ErrCommentOwner) {
		return true
	}
	if errors.Is(err, ErrFeedNotFound) {
		return true
	}
	if errors.Is(err, ErrFeedOwner) {
		return true
	}
	return false
}
