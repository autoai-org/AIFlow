// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)

// DefaultLogger is the instance shared by all modules
var DefaultLogger *logrus.Logger

// NewDefaultLogger returns the Logger Object
func NewDefaultLogger(logPath string) *logrus.Logger {
	if DefaultLogger != nil {
		return DefaultLogger
	}
	CreateFolderIfNotExist(filepath.Dir(logPath))
	writer, err := rotatelogs.New(
		logPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)
	if err != nil {
		logrus.Error("Cannot Initialize Logger..")
		logrus.Error(err.Error())
	}
	DefaultLogger = logrus.New()
	DefaultLogger.Hooks.Add(lfshook.NewHook(
		writer,
		&logrus.JSONFormatter{},
	))
	return DefaultLogger
}