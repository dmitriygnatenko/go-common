package logger

type EmailWriter struct {
}

func (w EmailWriter) Write(p []byte) (n int, err error) {
	return 0, err
}
