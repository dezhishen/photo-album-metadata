package inits

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/dezhishen/photo-album-metadata/pkg/config"
	"github.com/sirupsen/logrus"
)

func init() {
	//最大优先级 int8.Max
	registerWithPriority(doInitLogger, 0)
}

func defaultLoggerConfig(cfg *config.Config) *config.LoggerConfig {
	return &config.LoggerConfig{
		Output: "stdout",
		Level:  "info",
	}
}

func doInitLogger(cfg *config.Config) error {
	loggerConfig := cfg.LoggerConf
	if loggerConfig == nil {
		loggerConfig = defaultLoggerConfig(cfg)
		cfg.LoggerConf = loggerConfig
		err := config.Save(cfg)
		if err != nil {
			return err
		}
	}
	if loggerConfig.Output == "file" {
		if loggerConfig.OutputPath == "" {
			loggerConfig.OutputPath = "logs"
		}
		file, err := os.OpenFile(loggerConfig.OutputPath+string(filepath.Separator)+"app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		logrus.SetOutput(file)
	} else {
		logrus.SetOutput(os.Stdout)
	}
	switch strings.ToLower(loggerConfig.Level) {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetReportCaller(true)
	return nil
}
