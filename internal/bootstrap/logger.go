package bootstrap

import (
	"github.com/spleeroosh/messago/internal/config"
	"github.com/spleeroosh/messago/internal/pkg/application"
	"github.com/spleeroosh/messago/internal/pkg/logger"
)

func newLogger(conf config.Config, buildVersion application.BuildVersion) logger.Logger {
	return logger.NewLogger(
		conf.App.Name,
		logger.WithEnv(conf.App.Environment),
		logger.WithLevel(logger.Level(conf.App.LogLevel)),
		logger.WithBuildCommit(buildVersion.Commit),
		logger.WithBuildTime(buildVersion.Time),
		logger.WithPrettify(conf.App.PrettyLogs),
		logger.WithOverrideStdLogOut(true),
	)
}
