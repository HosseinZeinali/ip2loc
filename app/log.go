package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type LogEvent struct {
	id      int
	message string
}

const LOG_BASE_PATH = "var/log/"

var (
	invalidArgMessage      = LogEvent{id: 1, message: "Invalid arg: %s"}
	invalidArgValueMessage = LogEvent{id: 2, message: "Invalid value for argument: %s: %v"}
	missingArgMessage      = LogEvent{id: 3, message: "Missing arg: %s"}
	actionError            = LogEvent{id: 4, message: "error %s"}
)

func (m *MainLogger) InvalidArgMessage(argumentName string) {
	m.Errorf(invalidArgMessage.message, argumentName)
}

func (m *MainLogger) InvalidArgValue(argumentName string, argumentValue string) {
	m.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

func (m *MainLogger) MissingArg(argumentName string) {
	m.Errorf(missingArgMessage.message, argumentName)
}

func (m *MainLogger) ActionError(action string) {
	m.Errorf(actionError.message, action)
}

func (m *MainLogger) ActionInfo(action string) {
	m.Info(action)
}

type MainLogger struct {
	*logrus.Logger
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func NewLogger() *MainLogger {
	f, err := os.OpenFile(LOG_BASE_PATH+"logrus.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("Could Not Open Log File : " + err.Error())
	}
	var baseLogger = logrus.New()
	var standardLogger = &MainLogger{baseLogger}
	standardLogger.Formatter = &logrus.JSONFormatter{}
	mw := io.MultiWriter(os.Stdout, f)
	standardLogger.SetOutput(mw)
	return standardLogger
}
