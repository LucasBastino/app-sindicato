package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log = logrus.New()

func ConfigLogger() {

	// Lumberjack para rotacion de logs
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     31,
		Compress:   true,
	}

	// Creo un multiwriter para consola y archivo rotativo
	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)

	// Configuro para que su writer ea el multiwriter
	Log.SetOutput(multiWriter)

	// Uso foprmato JSON para los logs
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Nivel minimo de log
	Log.SetLevel(logrus.InfoLevel)
}