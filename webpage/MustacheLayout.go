/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
package webpage

import (
	"io/ioutil"
	"net/http"

	"github.com/cbroglie/mustache"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

/*
MustacheLayout is a layout using Mustache templates
*/
type MustacheLayout struct {
	layoutContents string
}

/*
LoadLayoutFile loads a layout from a file
*/
func (l *MustacheLayout) LoadLayoutFile(fileName string) error {
	var err error
	var layoutContents []byte

	if layoutContents, err = ioutil.ReadFile(fileName); err != nil {
		return err
	}

	l.layoutContents = string(layoutContents[:len(layoutContents)])
	return nil
}

/*
LoadLayoutString loads a layout from a passed-in byte array
*/
func (l *MustacheLayout) LoadLayoutString(contents []byte) error {
	l.layoutContents = string(contents[:len(contents)])
	return nil
}

/*
RenderViewFile renders a view from a file into this layout
*/
func (l *MustacheLayout) RenderViewFile(fileName string, context interface{}) (string, error) {
	var viewContents []byte
	var err error

	if viewContents, err = ioutil.ReadFile(fileName); err != nil {
		return "", errors.Wrapf(err, "Unable to read the view file %s", fileName)
	}

	return mustache.RenderInLayout(string(viewContents[:len(viewContents)]), l.layoutContents, context)
}

/*
RenderViewFilef renders a view from a file into this layout, then
writes it out to the provider writer. Useful for HTTP responses
*/
func (l *MustacheLayout) RenderViewFilef(ctx echo.Context, fileName string, context interface{}) error {
	var renderedContents string
	var err error

	if renderedContents, err = l.RenderViewFile(fileName, context); err != nil {
		return errors.Wrapf(err, "Unable to render the view file %s", fileName)
	}

	ctx.HTML(http.StatusOK, renderedContents)
	return nil
}

/*
RenderViewString renders a view from byte array content into this layout
*/
func (l *MustacheLayout) RenderViewString(contents []byte, context interface{}) (string, error) {
	return mustache.RenderInLayout(string(contents[:len(contents)]), l.layoutContents, contents)
}

/*
RenderViewStringf renders a view from byte array content into this layout, and
writes it out to the provided writer. Useful for HTTP responses
*/
func (l *MustacheLayout) RenderViewStringf(ctx echo.Context, contents []byte, context interface{}) error {
	var renderedContents string
	var err error

	if renderedContents, err = l.RenderViewString(contents, context); err != nil {
		return errors.Wrap(err, "Unable to render the view")
	}

	ctx.HTML(http.StatusOK, renderedContents)
	return nil
}
