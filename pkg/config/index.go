package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	RootPath   string `yaml:"root-path"`
	ListenAddr string `yaml:"listen-addr"`
}

var _config *Config
var rwlock sync.RWMutex

func Get() *Config {
	rwlock.RLock()
	defer rwlock.RUnlock()
	return _config
}

func Init(path string) error {
	rwlock.Lock()
	defer rwlock.Unlock()
	if _config != nil {
		return nil
	}
	cfg, err := parser(path)
	if err != nil {
		return err
	}
	_config = cfg
	return nil
}

func Reload(path string) error {
	rwlock.Lock()
	defer rwlock.Unlock()
	cfg, err := parser(path)
	if err != nil {
		return err
	}
	_config = cfg
	return nil
}

func parser(dir string) (*Config, error) {
	//如果路径不存在，则创建
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Println("mkdir dir failed")
			return nil, err
		}
	}
	log.Printf("config path: %s", dir)
	c := viper.New()
	c.SetConfigName("config")
	c.SetConfigType("yaml")
	c.AddConfigPath(dir)
	c.SetDefault("root-path", "/data")
	c.SetDefault("listen-addr", ":8080")
	if err := c.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file not found, use default config")
			//创建默认配置文件
			_, err = os.Create(dir + string(filepath.Separator) + "config.yaml")
			if err != nil {
				log.Println("create default config failed")
				return nil, err
			}
			err = c.WriteConfig()
			if err != nil {
				log.Println("write default config failed")
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	cfg := &Config{}
	err := c.Unmarshal(cfg)
	return cfg, err
}
