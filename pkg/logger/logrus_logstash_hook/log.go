package logrus_logstash_hook

import (
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

type LogrusToLogstashLOG struct {
	Log *logrus.Logger
}

var Logrus = LogrusToLogstashLOG{
	Log: connectLogstash(),
}

func connectLogstash() *logrus.Logger{
	err := godotenv.Load()
	if err != nil {
		return nil
	}

	log := logrus.New()
	conn, err := net.Dial("tcp", os.Getenv("LOGSTASH_HOST"))
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "ClientWorkerService"}))
	log.Hooks.Add(hook)
	return  log
}


func (l *LogrusToLogstashLOG) SendInfoLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Info(message)
}

func (l *LogrusToLogstashLOG)SendTraceLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Trace(message)
}

func (l *LogrusToLogstashLOG)SendDebugLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Debug(message)
}

func (l *LogrusToLogstashLOG)SendWarnLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Warn(message)
}

func (l *LogrusToLogstashLOG)SendErrorLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Error(message)
}

func (l *LogrusToLogstashLOG)SendFatalLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Fatal(message)
}

func (l *LogrusToLogstashLOG)SendPanicLog(parentStruct string, methodInfo string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Panic(message)
}

func (l *LogrusToLogstashLOG) SendInfofLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Infof(format,message)
}

func (l *LogrusToLogstashLOG)SendTracefLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Tracef(format, message)
}

func (l *LogrusToLogstashLOG)SendDebugfLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Debugf(format, message)
}

func (l *LogrusToLogstashLOG)SendWarnfLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Warnf(format, message)
}

func (l *LogrusToLogstashLOG)SendErrorfLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Errorf(format, message)
}

func (l *LogrusToLogstashLOG)SendFatalfLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Fatalf(format, message)
}

func (l *LogrusToLogstashLOG)SendPanicfLog(parentStruct string, methodInfo string, format string, message ...interface{}){
	ctx := l.Log.WithFields(logrus.Fields{
		"struct": parentStruct,
		"method": methodInfo,
	})
	ctx.Panicf(format, message)
}

