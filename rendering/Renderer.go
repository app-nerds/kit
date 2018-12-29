package rendering

import (
	"html/template"
	"io"
	"reflect"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
	logger    *logrus.Entry
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
func NewRenderer(logger *logrus.Entry) *Renderer {
	return &Renderer{
		logger:    logger,
		templates: make(map[string]*template.Template),
	}
}

func (r *Renderer) addHelperFunctions(t *template.Template) *template.Template {
	var funcMap = template.FuncMap{
		"eq": func(a, b interface{}) bool {
			return a == b
		},
		"neq": func(a, b interface{}) bool {
			return a != b
		},
		"arrayContainsString": func(array, value interface{}) bool {
			result := false
			iter := reflect.ValueOf(array)

			if iter.IsValid() {
				for i := 0; i < iter.Len(); i++ {
					if iter.Index(i).String() == value.(string) {
						result = true
						break
					}
				}
			}

			return result
		},
	}

	t.Funcs(funcMap)
	return t
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

	t = r.addHelperFunctions(t)
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
	err := r.templates[name].ExecuteTemplate(w, name, data)

	if err != nil {
		r.logger.WithError(err).WithFields(logrus.Fields{
			"name": name,
		}).Errorf("Error rendering template")

		return errors.Wrapf(err, "Error rendering template %s", name)
	}

	return nil
}
