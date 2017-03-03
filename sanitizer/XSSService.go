/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
package sanitizer

import "github.com/microcosm-cc/bluemonday"

/*
XSSService implements the XSSServiceProvider interface and offers functions to
help address cross-site script and sanitization concerns.
*/
type XSSService struct {
	sanitizer *bluemonday.Policy
}

/*
NewXSSService creates a new cross-site scripting service.
*/
func NewXSSService() *XSSService {
	policy := bluemonday.UGCPolicy()
	policy.AllowAttrs("align", "class", "style").OnElements("table", "div", "p", "section", "article", "header", "img", "span")
	policy.AllowAttrs("width", "height", "src", "frameborder", "allowfullscreen").OnElements("iframe")

	return &XSSService{
		sanitizer: policy,
	}
}

/*
SanitizeString attempts to sanitize a string by removing potentially dangerous
HTML/JS markup.
*/
func (service *XSSService) SanitizeString(input string) string {
	return service.sanitizer.Sanitize(input)
}
