# Captcha

This package provides services to add a captcha to your web applications. The following CAPTCHA services are supported.

* Google ReCAPTCHA v2

## Examples

### Google ReCAPTCHA v2

```golang
import "github.com/app-nerds/kit4/captcha"

captchaService := captcha.NewGoogleRecaptchaService(captcha.GoogleRecaptchaServiceConfig{
  CaptchaSecret: "secret",
})

if verifyCaptchaResponse, err := captchaService.VerifyCaptcha(captchaTokenFromFrontEnd, ipAddress); err != nil {
   // error
}

if !verifyCaptchaResponse.Success {
  // No bueno!
}
```

