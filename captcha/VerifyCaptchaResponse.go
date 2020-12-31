/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package captcha

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"
)

type VerifyCaptchaResponse struct {
	Success            bool      `json:"success"`
	ChallengeTimestamp time.Time `json:"challenge_ts"`
	HostName           string    `json:"hostname"`
	ErrorCodes         []string  `json:"error-codes"`
}

func NewVerifyCaptchaResponseFromReader(reader io.Reader) (VerifyCaptchaResponse, error) {
	var (
		err    error
		body   []byte
		result VerifyCaptchaResponse
	)

	if body, err = ioutil.ReadAll(reader); err != nil {
		return result, fmt.Errorf("error reading body when creating new VerifyCaptchaResponse struct: %w", err)
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("error unmarshaling JSON when creating new VerifyCaptchaResponse struct: %w", err)
	}

	return result, nil
}
