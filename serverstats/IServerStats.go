package serverstats

import "github.com/labstack/echo"

/*
IServerStats defines an interface for capturing and retrieving
server statistics
*/
type IServerStats interface {
	Handler(ctx echo.Context) error
	Middleware(next echo.HandlerFunc) echo.HandlerFunc
}
