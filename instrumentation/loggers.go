package instrumentation

type loggersType struct {
	instances []func(string)
}

type Loggers interface {
	Add(logger func(string))
	Log(message string)
}

type LogWriter interface {
	Log(message string)
}

func newLoggers() *loggersType {
	l := &loggersType{
		instances: []func(string){},
	}
	return l
}

func (l *loggersType) Add(logger func(string)) {
	l.instances = append(l.instances, logger)
}

func (l *loggersType) Log(message string) {
	for _, instance := range l.instances {
		instance(message)
	}
}
