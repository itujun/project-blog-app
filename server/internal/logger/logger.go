package logger

import "go.uber.org/zap"

// global var untuk singleton
var sugared *zap.SugaredLogger

// Init membuat logger produksi dengan Zap, panggil sekali di main()
func Init() {
	l, _ := zap.NewProduction()	// gunakan konfigurasi produksi
	sugared = l.Sugar()			// sugared: API logging yang ergonomis
}

// L mengembalikan instance singleton SugaredLogger
func L() *zap.SugaredLogger {
	if sugared == nil {
		Init()
	}
	return sugared
}
