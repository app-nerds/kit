/*
 * Copyright (c) 2021. App Nerds LLC All Rights Reserved
 */

package captcha

/*
CaptchaService describes methods for working with Google Captcha
*/
type CaptchaService interface {
	VerifyCaptcha(token string, ip string) (VerifyCaptchaResponse, error)
}
