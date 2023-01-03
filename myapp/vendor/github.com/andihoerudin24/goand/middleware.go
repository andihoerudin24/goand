package goand

import "net/http"

func (g *Goand) SessionLoad(next http.Handler) http.Handler {
	g.InfoLog.Println("SessionLoad called")
	return g.Session.LoadAndSave(next)
}
