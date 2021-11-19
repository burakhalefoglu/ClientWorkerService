package logger

type ILog interface {
	SendInfoLog(parentStruct string, methodInfo string, message ...interface{})
	SendTraceLog(parentStruct string, methodInfo string, message ...interface{})
	SendDebugLog(parentStruct string, methodInfo string, message ...interface{})
	SendWarnLog(parentStruct string, methodInfo string, message ...interface{})
	SendErrorLog(parentStruct string, methodInfo string, message ...interface{})
	SendFatalLog(parentStruct string, methodInfo string, message ...interface{})
	SendPanicLog(parentStruct string, methodInfo string, message ...interface{})
	SendInfofLog(parentStruct string, methodInfo string, format string, message ...interface{})
	SendTracefLog(parentStruct string, methodInfo string, format string, message ...interface{})
	SendDebugfLog(parentStruct string, methodInfo string, format string, message ...interface{})
	SendWarnfLog(parentStruct string, methodInfo string, format string, message ...interface{})
	SendErrorfLog(parentStruct string, methodInfo string, format string, message ...interface{})
	SendFatalfLog(parentStruct string, methodInfo string, format string, message ...interface{})
	SendPanicfLog(parentStruct string, methodInfo string, format string, message ...interface{})
}