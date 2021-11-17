package log

import (
	"context"
	"log"
	"net"
	"path"
	"runtime"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const MDC = "mdc"

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	// logrus.SetReportCaller(true)
	// logrus.SetFormatter(&prefixed.TextFormatter{
	// 	DisableColors:   true,
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// 	FullTimestamp:   true,
	// 	ForceFormatting: true,
	// })
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// logrus.SetFormatter(&logrus.JSONFormatter{CallerPrettyfier: func(f *runtime.Frame) (string, string) {
	// 	filename := path.Base(f.File)
	// 	return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
	// }})

	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "go-rest-crud"}))
	logrus.AddHook(hook)
}

func SetMdc(c *gin.Context, key string, value interface{}) {
	f := getMdc(c)
	if f == nil {
		f = logrus.Fields{}
	}
	f[key] = value
	c.Set(MDC, f)
}

func Info(c context.Context, args ...interface{}) {
	withFields(c).Info(args...)
}

func Infof(c context.Context, format string, args ...interface{}) {
	withFields(c).Infof(format, args...)
}

func Debug(c context.Context, args ...interface{}) {
	withFields(c).Debug(args...)
}

func Debugf(c context.Context, format string, args ...interface{}) {
	withFields(c).Debugf(format, args...)
}

func Error(c context.Context, args ...interface{}) {
	withFields(c).Error(args...)
}

func Errorf(c context.Context, format string, args ...interface{}) {
	withFields(c).Errorf(format, args...)
}

func withFields(c context.Context) *logrus.Entry {
	return logrus.WithFields(getMdc(c))
}

func getMdc(c context.Context) logrus.Fields {
	val := c.Value(MDC)
	var f logrus.Fields
	if val != nil {
		f = val.(logrus.Fields)
	} else {
		f = logrus.Fields{}
	}
	funcName, file, line := fileInfo(4)
	f["func"] = funcName
	f["file"] = file
	f["line"] = line
	return f
}

func fileInfo(skip int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(skip)
	var funcName string
	if !ok {
		file = "<???>"
		funcName = "<???>"
		line = 0
	} else {
		funcName = runtime.FuncForPC(pc).Name()
		file = path.Base(file)
		// slash := strings.LastIndex(file, "/")
		// if slash >= 0 {
		// 	file = file[slash+1:]
		// }
	}
	return funcName, file, line
}
