package log

func Info(args ...interface{}) {
	logger.Info(args)
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Error(args ...interface{})  {
	logger.Error(args)
}