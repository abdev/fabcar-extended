package grifts

import (
	"github.com/abdev/fabcar-extended/web_app/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
