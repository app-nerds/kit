/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package captcha

import (
	"encoding/json"
	"fmt"
)

type VerifyCaptchaRequest struct {
	Secret   string `json:"secret"`
	Token    string `json:"response"`
	RemoteIP string `json:"remoteip"`
}

func (r VerifyCaptchaRequest) ToJSON() []byte {
	result, _ := json.Marshal(&r)
	return result
}

func (r VerifyCaptchaRequest) ToQueryString() []byte {
	return []byte(fmt.Sprintf("secret=%s&response=%s&remoteip=%s", r.Secret, r.Token, r.RemoteIP))
}
