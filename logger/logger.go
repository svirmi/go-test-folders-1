package logger

var Log *zap.Logger

func init() {
	var err error
	Log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
}
