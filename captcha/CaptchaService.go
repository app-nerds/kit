/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package captcha

type CaptchaService interface {
	VerifyCaptcha(token string, ip string) (VerifyCaptchaResponse, error)
}
