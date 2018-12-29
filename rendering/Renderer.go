package rendering

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

/*
IRenderer defines an interface for a renderer
*/
type IRenderer interface {
	AddTemplateWithLayout(templateItem *TemplateItem) error
	AddTemplatesWithLayout(templateItems ...*TemplateItem) error
	Render(w io.Writer, name string, data interface{}, ctx echo.Context) error
}

/*
Renderer implements the Echo renderer and IRenderer interfaces
*/
type Renderer struct {
	templates map[string]*template.Template
}

/*
A TemplateItem describes a template to be rendered. This includes
the template's layout, name, and content
*/
type TemplateItem struct {
	LayoutContent string
	Name          string
	PageContent   string
}

/*
NewRenderer creates a new renderer component with an
empty set of templates
*/
func NewRenderer() *Renderer {
	return &Renderer{
		templates: make(map[string]*template.Template),
	}
}

/*
AddTemplateWithLayout adds a single template item to the template
cache
*/
func (r *Renderer) AddTemplateWithLayout(templateItem *TemplateItem) error {
	var err error

	t := template.New(templateItem.Name)

	if t, err = t.Parse(templateItem.LayoutContent); err != nil {
		return errors.Wrapf(err, "Error parsing layout while attempting to add template %s", templateItem.Name)
	}

	if t, err = t.Parse(templateItem.PageContent); err != nil {
		return errors.Wrapf(err, "Error parsing page while attempting to add template %s", templateItem.Name)
	}

	r.templates[templateItem.Name] = t
	return nil
}

/*
AddTemplatesWithLayout adds one ore more template items
*/
func (r *Renderer) AddTemplatesWithLayout(templateItems ...*TemplateItem) error {
	var err error

	for _, templateItem := range templateItems {
		if err = r.AddTemplateWithLayout(templateItem); err != nil {
			return err
		}
	}

	return nil
}

/*
Render renders a template by name to the supplied writer
*/
func (r *Renderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	err := r.templates[name].Execute(w, data)

	if err != nil {
		fmt.Printf("\nError rendering template %s: %s\n", name, err.Error())
		return errors.Wrapf(err, "Error rendering template %s", name)
	}

	return nil
}
