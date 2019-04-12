package syslogrus

import (
	"fmt"
	"log/syslog"
	"os"

	"github.com/sirupsen/logrus"
)

type SyslogHookConfig struct {
	Formatter logrus.Formatter
	Network   string
	Priority  syslog.Priority
	Raddr     string
	Tag       string
}

// SyslogHook to send logs via syslog.
type SyslogHook struct {
	Formatter     logrus.Formatter
	SyslogNetwork string
	SyslogRaddr   string
	Writer        *syslog.Writer
}

// Creates a hook to be added to an instance of logger.
//
// `hook, err := NewSyslogHook(SyslogHookConfig{Priority: syslog.LOG_DEBUG})`
// `if err == nil { log.Hooks.Add(hook) }`
func NewSyslogHook(config SyslogHookConfig) (*SyslogHook, error) {
	w, err := syslog.Dial(config.Network, config.Raddr, config.Priority, config.Tag)
	return &SyslogHook{config.Formatter, config.Network, config.Raddr, w}, err
}

func (hook *SyslogHook) Fire(entry *logrus.Entry) error {
	var line string
	var err error
	if hook.Formatter != nil {
		bytes, err := hook.Formatter.Format(entry)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to format entry, %v", err)
			return err
		}
		line = string(bytes)
	} else {
		line, err = entry.String()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
			return err
		}
	}

	switch entry.Level {
	case logrus.PanicLevel:
		return hook.Writer.Crit(line)
	case logrus.FatalLevel:
		return hook.Writer.Crit(line)
	case logrus.ErrorLevel:
		return hook.Writer.Err(line)
	case logrus.WarnLevel:
		return hook.Writer.Warning(line)
	case logrus.InfoLevel:
		return hook.Writer.Info(line)
	case logrus.DebugLevel, logrus.TraceLevel:
		return hook.Writer.Debug(line)
	default:
		return nil
	}
}

func (hook *SyslogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
