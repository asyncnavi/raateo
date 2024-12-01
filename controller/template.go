package controller

import (
	"github.com/asyncnavi/raateo/database"
)

type Template struct {
	Title        string
	User         *database.User
	FieldErrors  map[string]string
	ErrorMessage string
}

func NewTemplate(title string) *Template {
	return &Template{
		Title:       title,
		FieldErrors: map[string]string{},
	}
}

func (t *Template) AddFieldError(field string, message string) {
	t.FieldErrors[field] = message
}

func (t *Template) AddErrorMessage(message string) {
	t.ErrorMessage = message
}

func (t *Template) HasErrors() bool {
	return len(t.FieldErrors) != 0
}

func (t *Template) AddTitle(title string) {
	t.Title = title
}

func (t *Template) AddUser(user *database.User) {
	t.User = user
}
