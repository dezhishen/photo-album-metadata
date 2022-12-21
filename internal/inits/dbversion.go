package inits

import (
	"github.com/dezhishen/photo-album-metadata/internal/dbversion"
	"github.com/dezhishen/photo-album-metadata/pkg/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	register(doInitDB)
}
func doInitDB(cfg *config.Config) error {
	log.Info("start init database version")
	return dbversion.StartVersion(cfg.MetadataPath)
}
