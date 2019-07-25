package customcontext_test

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/app-nerds/kit/v4/customcontext"
	"github.com/labstack/echo/v4"
)

func TestGetAdminUserContext(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	ctx := e.NewContext(request, response)

	expected := &customcontext.AdminUserContext{
		Context:  ctx,
		UserID:   "user",
		UserName: "Bob",
	}

	actual := customcontext.GetAdminUserContext(expected)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %+v to be %+v", actual, expected)
	}
}

func TestGetAdminUserContext_MakesNewAdminUserContext(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	ctx := e.NewContext(request, response)

	expected := &customcontext.AdminUserContext{
		Context: ctx,
	}

	actual := customcontext.GetAdminUserContext(ctx)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %+v to be %+v", actual, expected)
	}
}
