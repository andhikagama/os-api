package resource

import (
	"github.com/sirupsen/logrus"
)

// NewLogger .
func NewLogger() *logrus.Logger {
	return logrus.New()
}
