/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package rendering

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"reflect"

	"github.com/app-nerds/kit/v5/datetime"
	"github.com/app-nerds/kit/v5/sanitizer"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

/*
IRenderer defines an interface for a renderer
*/
type IRenderer interface {
	AddTemplates(templateItems ...*Template) error
	AddTemplatesWithLayout(templateItems ...*TemplateWithLayout) error
	LoadFile(fileName string) string
	Render(w io.Writer, name string, data interface{}, ctx echo.Context) error
	SetDebug(value bool)
}

/*
Renderer implements the Echo renderer and IRenderer interfaces
*/
type Renderer struct {
	dateTimeParser datetime.IDateTimeParser
	debug          bool
	logger         *logrus.Entry
	templates      map[string]*template.Template
	xssSanitizer   sanitizer.IXSSServiceProvider
}

/*
A Template is a simple Go template with a name
*/
type Template struct {
	Name        string
	PageContent string
}

/*
A TemplateWithLayout describes a template to be rendered. This includes
the template's layout, name, and content
*/
type TemplateWithLayout struct {
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
		dateTimeParser: &datetime.DateTimeParser{},
		debug:          false,
		logger:         logger,
		templates:      make(map[string]*template.Template),
		xssSanitizer:   sanitizer.NewXSSService(),
	}
}

func (r *Renderer) addHelperFunctions(t *template.Template) *template.Template {
	var funcMap = template.FuncMap{
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
		"toUSDateTime": r.dateTimeParser.ToUSDateTime,
		"toUSDate":     r.dateTimeParser.ToUSDate,
		"toUSTime":     r.dateTimeParser.ToUSTime,
		"toSQLString":  r.dateTimeParser.ToSQLString,
		"toISO8601":    r.dateTimeParser.ToISO8601,
		"sanitize":     r.xssSanitizer.SanitizeString,
		"intLT": func(value1, value2 int) bool {
			return value1 < value2
		},
		"intLTE": func(value1, value2 int) bool {
			return value1 <= value2
		},
		"intGT": func(value1, value2 int) bool {
			return value1 > value2
		},
		"intGTE": func(value1, value2 int) bool {
			return value1 >= value2
		},
		"neq": func(a, b interface{}) bool {
			return a != b
		},
	}

	t = t.Funcs(funcMap)
	return t
}

/*
addTemplateWithLayout adds a single template item to the template
cache
*/
func (r *Renderer) addTemplate(templateItem *Template) error {
	var err error

	if r.debug {
		fmt.Printf("Adding %s...\n", templateItem.Name)
	}

	t := template.New(templateItem.Name)
	t = r.addHelperFunctions(t)

	if t, err = t.Parse(templateItem.PageContent); err != nil {
		if r.debug {
			fmt.Printf("\tError parsing template - %s\n", err.Error())
		}

		return errors.Wrapf(err, "Error parsing page while attempting to add template %s", templateItem.Name)
	}

	r.templates[templateItem.Name] = t

	if r.debug {
		fmt.Printf("\tTemplate parsed successfully!\n")
	}

	return nil
}

/*
addTemplateWithLayout adds a single template item with a layout to the template
cache
*/
func (r *Renderer) addTemplateWithLayout(templateItem *TemplateWithLayout) error {
	var err error

	if r.debug {
		fmt.Printf("Adding %s...\n", templateItem.Name)
	}

	t := template.New(templateItem.Name)
	t = r.addHelperFunctions(t)

	if t, err = t.Parse(templateItem.LayoutContent); err != nil {
		if r.debug {
			fmt.Printf("\tError parsing layout - %s\n", err.Error())
		}

		return errors.Wrapf(err, "Error parsing layout while attempting to add template %s", templateItem.Name)
	}

	if t, err = t.Parse(templateItem.PageContent); err != nil {
		panic("Error parsing template '" + templateItem.Name + "' - " + err.Error())
	}

	r.templates[templateItem.Name] = t

	if r.debug {
		fmt.Printf("\tTemplate parsed successfully!\n")
	}

	return nil
}

/*
AddTemplates adds one or more templates
*/
func (r *Renderer) AddTemplates(templateItems ...*Template) error {
	var err error

	if r.debug {
		fmt.Printf("\n\n-------------------------------------------------------------------------------\n")
		fmt.Printf("\nAdding %d templates...\n", len(templateItems))
	}

	for _, templateItem := range templateItems {
		if err = r.addTemplate(templateItem); err != nil {
			return err
		}
	}

	if r.debug {
		fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
	}

	return nil
}

/*
AddTemplatesWithLayout adds one ore more templates with layouts
*/
func (r *Renderer) AddTemplatesWithLayout(templateItems ...*TemplateWithLayout) error {
	var err error

	if r.debug {
		fmt.Printf("\n\n-------------------------------------------------------------------------------\n")
		fmt.Printf("\nAdding %d templates with layouts...\n", len(templateItems))
	}

	for _, templateItem := range templateItems {
		if err = r.addTemplateWithLayout(templateItem); err != nil {
			return err
		}
	}

	if r.debug {
		fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
	}

	return nil
}

/*
LoadFile reads all the contents of a file
*/
func (r *Renderer) LoadFile(fileName string) string {
	var err error
	var result []byte

	if result, err = ioutil.ReadFile(fileName); err != nil {
		if r.debug {
			fmt.Printf("Error reading file %s in LoadFile", fileName)
		}

		return ""
	}

	return string(result)
}

/*
Render renders a template by name to the supplied writer
*/
func (r *Renderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	var ok bool

	if _, ok = r.templates[name]; !ok {
		if r.debug {
			fmt.Printf("\nTemplate %s not found\n", name)
		}

		return errors.New("Template " + name + " not found")
	}

	err := r.templates[name].ExecuteTemplate(w, name, data)

	if err != nil {
		r.logger.WithError(err).WithFields(logrus.Fields{
			"name": name,
		}).Errorf("Error rendering template")

		return errors.Wrapf(err, "Error rendering template %s", name)
	}

	return nil
}

func (r *Renderer) SetDebug(value bool) {
	r.debug = value
}
