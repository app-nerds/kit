/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package logging2

import (
	fireplacehook "github.com/app-nerds/fireplace/cmd/fireplace-hook"
	"github.com/sirupsen/logrus"
)

/*
NewFireplaceLogger creates a new logrus logger with a Fireplace Server
hook installed
*/
func NewFireplaceLogger(application, logLevel, fireplaceURL string, fields logrus.Fields) *logrus.Entry {
	var err error
	var level logrus.Level

	if level, err = logrus.ParseLevel(logLevel); err != nil {
		level = logrus.ErrorLevel
	}

	logger := logrus.New().WithFields(fields)
	logger.Logger.SetLevel(level)

	logger.Logger.AddHook(fireplacehook.NewFireplaceHook(&fireplacehook.FireplaceHookConfig{
		Application:  application,
		FireplaceURL: fireplaceURL,
	}))

	return logger
}
