# Sanitizer

This package provides a server for sanitizing strings by removing potentially dangerous HTML and JavaScript that may be used to execute XSS attacks.

```go
sampleHTML := `<h1>This is a test</h1>
<script>
	alert("Hi");
</script>`

xssService := sanitizer.NewXSSService()
sanitizedHTML := xssService.SanitizeString(sampleHTML)

// New HTML == `<h1>This is a test</h1>`
```
