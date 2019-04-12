# Syslogrus 

Syslogrus is a Logrus hook with support of formatting.

## Usage

```go
import (
  "log/syslog"
  "github.com/sirupsen/logrus"
  syslogrus "github.com/magicsandbox/syslogrus"
)

func main() {
	log := logrus.New()
	hook, _ := NewSyslogHook(
		SyslogHookConfig{
			Network:  "udp",
			Raddr:    "localhost:514",
			Priority: syslog.LOG_INFO,
			Formatter: &logrus.JSONFormatter{},
		})

	log.Hooks.Add(hook)
}
```

If you want to connect to local syslog, just leave `Network` and `Raddr` empty. It should look like this:

```go
import (
  "log/syslog"
  "github.com/sirupsen/logrus"
  syslogrus "github.com/magicsandbox/syslogrus"
)

func main() {
	log := logrus.New()
	hook, _ := NewSyslogHook(SyslogHookConfig{Priority: syslog.LOG_INFO})
	log.Hooks.Add(hook)
}
```
