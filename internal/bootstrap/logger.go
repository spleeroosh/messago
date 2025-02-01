package bootstrap

import (
	"github.com/spleeroosh/messago/internal/config"
	"gitlab.xbet.lan/web-backend/go/pkg/log"

	"github.com/spleeroosh/messago/internal/infrastructure/http/application"
	"github.com/spleeroosh/messago/internal/infrastructure/logger"
)

func newLogger(conf config.Config, buildVersion application.BuildVersion) logger.Logger {
	return log.NewLogger(
		conf.App.Name,
		log.WithEnv(conf.App.Environment),
		log.WithLevel(log.Level(conf.App.LogLevel)),
		log.WithBuildCommit(buildVersion.Commit),
		log.WithBuildTime(buildVersion.Time),
		log.WithPrettify(conf.App.PrettyLogs),
		log.WithOverrideStdLogOut(true),
	)
}
