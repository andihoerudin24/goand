package middleware

import (
	"github.com/andihoerudin24/goand"
	"myapp/data"
)

type Middleware struct {
	App    *goand.Goand
	Models data.Models
}
