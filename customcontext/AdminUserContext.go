// Copyright 2013-2018 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package customcontext

import "github.com/labstack/echo"

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
