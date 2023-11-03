package instrumentation

var loggers Loggers

func init() {
	loggers = newLoggers()
}

func AddLogger(logger func(string)) {
	loggers.Add(logger)
}

func Log(message string) {
	loggers.Log(message)
}
