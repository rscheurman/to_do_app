package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/rscheurman/ToDoApp/to_do_app/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
