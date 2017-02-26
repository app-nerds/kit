/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
package webpage

import "github.com/labstack/echo"

/*
ILayout defines an interface for layouts to adhere to
*/
type ILayout interface {
	LoadLayoutFile(fileName string) error
	LoadLayoutString(contents []byte) error
	RenderViewFile(fileName string, context interface{}) (string, error)
	RenderViewFilef(ctx echo.Context, fileName string, context interface{}) error
	RenderViewString(contents []byte, context interface{}) (string, error)
	RenderViewStringf(ctx echo.Context, contents []byte, context interface{}) error
}
