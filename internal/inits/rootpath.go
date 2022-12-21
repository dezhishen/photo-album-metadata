package inits

import "github.com/dezhishen/photo-album-metadata/pkg/config"

func init() {
	register(InitRootPath)
}
func InitRootPath(cfg *config.Config) error {
	return nil
}
