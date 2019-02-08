package rendering

import (
	"io"

	"github.com/labstack/echo"
)

type MockRenderer struct {
	AddTemplatesFunc           func(templateItems ...*Template) error
	AddTemplatesWithLayoutFunc func(templateItems ...*TemplateWithLayout) error
	LoadFileFunc               func(fileName string) string
	RenderFunc                 func(w io.Writer, name string, data interface{}, ctx echo.Context) error
	SetDebugFunc               func(value bool)
}

func (m *MockRenderer) AddTemplates(templateItems ...*Template) error {
	return m.AddTemplatesFunc(templateItems...)
}

func (m *MockRenderer) AddTemplatesWithLayout(templateItems ...*TemplateWithLayout) error {
	return m.AddTemplatesWithLayoutFunc(templateItems...)
}

func (m *MockRenderer) LoadFile(fileName string) string {
	return m.LoadFileFunc(fileName)
}

func (m *MockRenderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	return m.RenderFunc(w, name, data, ctx)
}

func (m *MockRenderer) SetDebug(value bool) {
	m.SetDebugFunc(value)
}
