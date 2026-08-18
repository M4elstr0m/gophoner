package logs

import (
	"io"

	"github.com/M4elstr0m/gophoner/internal/version"
	"github.com/adrg/xdg"
	"github.com/charmbracelet/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	fileName  string = "gophoner.log"
	maxSizeMB int    = 10
	maxAgeDay int    = 7
)

func Init(noLog bool) error {
	if noLog {
		log.SetOutput(io.Discard)
		return nil
	}

	path, err := xdg.StateFile(version.APP_NAME + "/" + fileName)
	if err != nil {
		return err
	}

	log.SetOutput(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSizeMB,
		MaxBackups: 5,
		MaxAge:     maxAgeDay,
		Compress:   true,
	})
	log.SetFormatter(log.LogfmtFormatter)

	log.Info("gophoner started")
	return nil
}
