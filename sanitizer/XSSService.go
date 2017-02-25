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
	policy.AllowAttrs("align", "class", "style").OnElements("table", "div", "p", "section", "article", "header", "img")
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
