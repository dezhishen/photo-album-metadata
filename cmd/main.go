package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dezhishen/photo-album-metadata/internal/inits"
	"github.com/dezhishen/photo-album-metadata/internal/routes"
	"github.com/dezhishen/photo-album-metadata/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// log.Printf("current working directory: %s", wd)
	cfgPath := fmt.Sprintf("%s%sconfig", wd, string(filepath.Separator))
	// 初始化配置
	err = config.Init(cfgPath)
	if err != nil {
		panic(err)
	}
	cfg := config.GetConfig()
	doInit(cfg)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	routes.Init(r)
	r.Run(cfg.ListenAddr)
}

func doInit(cfg *config.Config) {
	jobs := inits.GetAll()
	if len(jobs) == 0 {
		return
	}
	for _, job := range jobs {
		handler := job
		err := handler(cfg)
		if err != nil {
			panic(err)
		}
	}
}
