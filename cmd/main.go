package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dezhishen/photo-album-metadata/internal/routes"
	"github.com/dezhishen/photo-album-metadata/pkg/config"
	"github.com/dezhishen/photo-album-metadata/pkg/fileutil"
	"github.com/gin-gonic/gin"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Printf("current working directory: %s", wd)
	err = config.Init(fmt.Sprintf("%s%sconfig", wd, string(filepath.Separator)))
	if err != nil {
		panic(err)
	}
	cfg := config.Get()
	err = fileutil.ScanFile(cfg.RootPath, func(line string) error { return nil })
	if err != nil {
		return
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	err = routes.Init(r)
	if err != nil {
		panic(err)
	}
	r.Run(cfg.ListenAddr)
}
