package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

func (e *Event) SetEvent(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}

func (e *Event) Event() string {
	return e.title
}
