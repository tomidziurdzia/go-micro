package todo

import "strings"

type Title string

func NewTitle(raw string) (Title, error) {
	s := strings.TrimSpace(raw)
	if s == "" {
		return "", ErrInvalidTitle
	}
	if len([]rune(s)) > 140 {
		return "", ErrInvalidTitle
	}
	return Title(s), nil
}
