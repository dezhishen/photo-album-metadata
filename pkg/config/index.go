package config

import (
	"os"
	"path/filepath"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var configInstance *configInstanceStruct
var gloabLock sync.RWMutex

type configInstanceStruct struct {
	sync.RWMutex
	cfg           *Config
	viperInstance *viper.Viper
}

type Config struct {
	RootPath     string        `yaml:"root-path"`
	ListenAddr   string        `yaml:"listen-addr"`
	MetadataPath string        `yaml:"metadata-path"`
	LoggerConf   *LoggerConfig `yaml:"logger"`
}

type LoggerConfig struct {
	// 级别
	Level string `yaml:"level"`
	// 输出方式,stdout/file
	Output string `yaml:"output"`
	// 输出文件路径
	OutputPath string `yaml:"output-path"`
}

func (_cfg *configInstanceStruct) SaveItem(key string, value interface{}) error {
	_cfg.Lock()
	defer _cfg.Unlock()
	_cfg.viperInstance.Set(key, value)
	err := _cfg.viperInstance.WriteConfig()
	if err != nil {
		return err
	}
	err = _cfg.viperInstance.ReadInConfig()
	if err != nil {
		return err
	}
	return _cfg.viperInstance.Unmarshal(_cfg.cfg)
}

func (_cfg *configInstanceStruct) Save(cfg *Config) error {
	_cfg.Lock()
	defer _cfg.Unlock()
	// _cfg.viperInstance.Set(key, value)
	// _cfg.viperInstance.Set()
	err := _cfg.viperInstance.WriteConfig()
	if err != nil {
		return err
	}
	err = _cfg.viperInstance.ReadInConfig()
	if err != nil {
		return err
	}
	return _cfg.viperInstance.Unmarshal(_cfg.cfg)
}

var _config *configInstanceStruct

func GetConfig() *Config {
	gloabLock.RLock()
	defer gloabLock.RUnlock()
	return _config.cfg
}

func SaveItem(key string, value interface{}) error {
	return _config.SaveItem(key, value)
}

func Save(cfg *Config) error {
	return _config.Save(cfg)
}

func Init(path string) error {
	gloabLock.Lock()
	defer gloabLock.Unlock()
	var err error
	configInstance, err = parse(path)
	_config = configInstance
	return err
}

func Reload(path string) error {
	gloabLock.Lock()
	defer gloabLock.Unlock()
	var err error
	configInstance, err = parse(path)
	return err
}

func getViperInstance(dir string) *viper.Viper {
	c := viper.New()
	c.SetConfigName("config")
	c.SetConfigType("yaml")
	c.AddConfigPath(dir)
	return c
}

func parse(dir string) (*configInstanceStruct, error) {
	//如果路径不存在，则创建
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Println("mkdir dir failed")
			return nil, err
		}
	}
	log.Printf("config path: %s", dir)
	c := getViperInstance(dir)
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
	if err != nil {
		return nil, err
	}
	return &configInstanceStruct{
		cfg:           cfg,
		viperInstance: c,
	}, nil
}
