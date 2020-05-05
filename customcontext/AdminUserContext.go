/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package customcontext

import "github.com/labstack/echo/v4"

type AdminUserContext struct {
	echo.Context
	UserID   string
	UserName string
}

/*
GetAdminUserContext takes an Echo context and coerces it to an
AdminUserContext
*/
func GetAdminUserContext(ctx echo.Context) *AdminUserContext {
	var ok bool

	if _, ok = ctx.(*AdminUserContext); ok {
		return ctx.(*AdminUserContext)
	}

	return &AdminUserContext{
		Context: ctx,
	}
}
