/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package serverstats

import "github.com/labstack/echo/v4"

/*
IServerStats defines an interface for capturing and retrieving
server statistics
*/
type IServerStats interface {
	GetAverageResponseTimeGraph(precision ResponseTimePrecision) ResponseTimeGraphCollection
	Handler(ctx echo.Context) error
	Middleware(next echo.HandlerFunc) echo.HandlerFunc
}
