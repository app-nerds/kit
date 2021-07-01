/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package captcha

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/app-nerds/kit/v5/restclient"
)

type GoogleRecaptchaServiceConfig struct {
	CaptchaSecret string
}

type GoogleRecaptchaService struct {
	CaptchaSecret string
	HttpClient    restclient.IHTTPClient
}

/*
NewGoogleRecaptchaService creates a new Captcha service that uses
Google Recaptcha
*/
func NewGoogleRecaptchaService(config GoogleRecaptchaServiceConfig) *GoogleRecaptchaService {
	return &GoogleRecaptchaService{
		CaptchaSecret: config.CaptchaSecret,
		HttpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

/*
VerifyCaptcha verifies the captcha request with the provider and returns a response
*/
func (s *GoogleRecaptchaService) VerifyCaptcha(token string, ip string) (VerifyCaptchaResponse, error) {
	var (
		err      error
		result   VerifyCaptchaResponse
		request  *http.Request
		response *http.Response
	)

	verifyRequest := VerifyCaptchaRequest{
		Secret:   s.CaptchaSecret,
		Token:    token,
		RemoteIP: ip,
	}

	if request, err = http.NewRequest(http.MethodPost, "https://www.google.com/recaptcha/api/siteverify", bytes.NewBuffer(verifyRequest.ToQueryString())); err != nil {
		return result, fmt.Errorf("error creating request to verify captcha: %w", err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if response, err = s.HttpClient.Do(request); err != nil {
		return result, fmt.Errorf("error making request to verify captcha: %w", err)
	}

	defer response.Body.Close()

	if result, err = NewVerifyCaptchaResponseFromReader(response.Body); err != nil {
		return result, fmt.Errorf("error creating response: %w", err)
	}

	return result, nil
}
