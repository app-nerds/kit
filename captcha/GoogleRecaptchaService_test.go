/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package captcha_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/app-nerds/kit/v5/captcha"
	"github.com/app-nerds/kit/v5/restclient"
)

func TestNewGoogleRecaptchaService(t *testing.T) {
	actual := captcha.NewGoogleRecaptchaService(captcha.GoogleRecaptchaServiceConfig{
		CaptchaSecret: "abc",
	})

	isTest := func(t interface{}) bool {
		switch t.(type) {
		case captcha.CaptchaService:
			return true
		default:
			return false
		}
	}

	if !isTest(actual) {
		t.Errorf("Expected object to be of type CaptchaService")

	}
}

func TestGoogleRecaptchaService_VerifyCaptcha(t *testing.T) {
	var (
		capturedURL string
	)

	type fields struct {
		captchaSecret string
		httpClient    restclient.IHTTPClient
	}

	type args struct {
		token string
		ip    string
	}

	successResult := captcha.VerifyCaptchaResponse{
		Success:    true,
		HostName:   "host",
		ErrorCodes: nil,
	}

	errorResult := map[string]string{
		"bad": "juju",
	}

	/*
	 * Setup tests
	 */
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    captcha.VerifyCaptchaResponse
		wantURL string
		wantErr bool
	}{
		{
			name: "Returns a response upon success",
			fields: fields{
				captchaSecret: "secret",
				httpClient: &restclient.MockHttpClient{
					DoFunc: func(req *http.Request) (*http.Response, error) {
						capturedURL = req.URL.String()

						resultRaw, _ := json.Marshal(&successResult)
						resultBody := bytes.NewReader(resultRaw)

						response := &http.Response{
							Status:     "200 OK",
							StatusCode: http.StatusOK,
							Body:       ioutil.NopCloser(resultBody),
						}

						return response, nil
					},
				},
			},
			args: args{
				token: "abcdefg",
				ip:    "::1",
			},
			want:    successResult,
			wantURL: "https://www.google.com/recaptcha/api/siteverify",
			wantErr: false,
		},
		{
			name: "Returns an error when the HTTP request fails",
			fields: fields{
				captchaSecret: "secret",
				httpClient: &restclient.MockHttpClient{
					DoFunc: func(req *http.Request) (*http.Response, error) {
						capturedURL = req.URL.String()

						return nil, fmt.Errorf("error in HTTP")
					},
				},
			},
			args: args{
				token: "abcdefg",
				ip:    "::1",
			},
			want:    captcha.VerifyCaptchaResponse{},
			wantURL: "https://www.google.com/recaptcha/api/siteverify",
			wantErr: true,
		},
		{
			name: "Returns an error when there is a problem deserializing the response",
			fields: fields{
				captchaSecret: "secret",
				httpClient: &restclient.MockHttpClient{
					DoFunc: func(req *http.Request) (*http.Response, error) {
						capturedURL = req.URL.String()

						resultRaw, _ := json.Marshal(&errorResult)
						resultBody := bytes.NewReader(resultRaw)

						response := &http.Response{
							Status:     "200 OK",
							StatusCode: http.StatusOK,
							Body:       ioutil.NopCloser(resultBody),
						}

						return response, fmt.Errorf("error in HTTP")
					},
				},
			},
			args: args{
				token: "abcdefg",
				ip:    "::1",
			},
			want:    captcha.VerifyCaptchaResponse{},
			wantURL: "https://www.google.com/recaptcha/api/siteverify",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &captcha.GoogleRecaptchaService{
				CaptchaSecret: tt.fields.captchaSecret,
				HttpClient:    tt.fields.httpClient,
			}

			got, err := service.VerifyCaptcha(tt.args.token, tt.args.ip)

			if tt.wantErr && err == nil {
				t.Errorf("TestGoogleRecaptchaService_VerifyCaptcha() wanted error, got nil")
			}

			if !tt.wantErr {
				if capturedURL != tt.wantURL {
					t.Errorf("TestGoogleRecaptchaService_VerifyCaptcha()\nwanted URL: %s\ngot URL: %s", tt.wantURL, capturedURL)
				}

				if !reflect.DeepEqual(tt.want, got) {
					t.Errorf("TestGoogleRecaptcha_VerifyCaptcha()\nwanted:\n%v\ngot:\n%v", tt.want, got)
				}
			}
		})
	}
}
