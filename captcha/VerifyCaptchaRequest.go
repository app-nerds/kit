/*
 * Copyright (c) 2021. App Nerds LLC All Rights Reserved
 */

package captcha

import (
	"encoding/json"
	"fmt"
)

/*
VerifyCaptchaRequest is used to request a Captcha verification
*/
type VerifyCaptchaRequest struct {
	Secret   string `json:"secret"`
	Token    string `json:"response"`
	RemoteIP string `json:"remoteip"`
}

/*
ToJSON converts this VerifyCaptchaRequest to JSON
*/
func (r VerifyCaptchaRequest) ToJSON() []byte {
	result, _ := json.Marshal(&r)
	return result
}

/*
ToQueryString returns a query string from this request's parameters
*/
func (r VerifyCaptchaRequest) ToQueryString() []byte {
	return []byte(fmt.Sprintf("secret=%s&response=%s&remoteip=%s", r.Secret, r.Token, r.RemoteIP))
}
