/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
package webpage

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

/*
GoLayout is a layout using Go templates
*/
type GoLayout struct {
	layout         *template.Template
	layoutContents string
}

func (l *GoLayout) addHelperFunctions() {
	var funcMap = template.FuncMap{
		"eq": func(a, b interface{}) bool {
			return a == b
		},
		"neq": func(a, b interface{}) bool {
			return a != b
		},
	}

	l.layout.Funcs(funcMap)
}

/*
LoadLayoutFile loads a layout from a file
*/
func (l *GoLayout) LoadLayoutFile(fileName string) error {
	var err error
	var layoutContents []byte

	if layoutContents, err = ioutil.ReadFile(fileName); err != nil {
		return err
	}

	l.layoutContents = string(layoutContents[:len(layoutContents)])
	l.renderLayout()
	return nil
}

/*
LoadLayoutString loads a layout from a passed-in byte array
*/
func (l *GoLayout) LoadLayoutString(contents []byte) error {
	l.layoutContents = string(contents[:len(contents)])
	l.renderLayout()
	return nil
}

func (l *GoLayout) renderLayout() {
	l.layout = template.Must(template.New("layout").Parse(l.layoutContents))
	l.addHelperFunctions()
}

/*
RenderViewFile renders a view from a file into this layout
*/
func (l *GoLayout) RenderViewFile(fileName string, context interface{}) (string, error) {
	var viewContents []byte
	var err error
	var viewTemplate *template.Template
	stringWriter := bytes.NewBufferString("")

	l.renderLayout()

	if viewContents, err = ioutil.ReadFile(fileName); err != nil {
		return "", errors.Wrapf(err, "Unable to read the view file %s", fileName)
	}

	if viewTemplate, err = l.layout.Parse(string(viewContents[:len(viewContents)])); err != nil {
		return "", errors.Wrapf(err, "Unable to parse the view file %s", fileName)
	}

	if err = viewTemplate.Execute(stringWriter, context); err != nil {
		return "", errors.Wrapf(err, "Unable to render the file %s", fileName)
	}

	return stringWriter.String(), nil
}

/*
RenderViewFilef renders a view from a file into this layout, then
writes it out to the provider writer. Useful for HTTP responses
*/
func (l *GoLayout) RenderViewFilef(ctx echo.Context, fileName string, context interface{}) error {
	var renderedContents string
	var err error

	l.renderLayout()

	if renderedContents, err = l.RenderViewFile(fileName, context); err != nil {
		return errors.Wrapf(err, "Unable to render the view file %s", fileName)
	}

	ctx.String(http.StatusOK, renderedContents)
	return nil
}

/*
RenderViewString renders a view from byte array content into this layout
*/
func (l *GoLayout) RenderViewString(contents []byte, context interface{}) (string, error) {
	var err error
	var viewTemplate *template.Template
	stringWriter := bytes.NewBufferString("")

	l.renderLayout()

	if viewTemplate, err = l.layout.Parse(string(contents[:len(contents)])); err != nil {
		return "", errors.Wrapf(err, "Unable to parse the view in RenderViewString()")
	}

	if err = viewTemplate.Execute(stringWriter, context); err != nil {
		return "", errors.Wrapf(err, "Unable to render view in RenderViewString()")
	}

	return stringWriter.String(), nil
}

/*
RenderViewStringf renders a view from byte array content into this layout, and
writes it out to the provided writer. Useful for HTTP responses
*/
func (l *GoLayout) RenderViewStringf(ctx echo.Context, contents []byte, context interface{}) error {
	var renderedContents string
	var err error

	l.renderLayout()

	if renderedContents, err = l.RenderViewString(contents, context); err != nil {
		return errors.Wrap(err, "Unable to render the view")
	}

	ctx.String(http.StatusOK, renderedContents)
	return nil
}
