package config

import (
	"github.com/jinzhu/configor"
	"os"
)

type Config struct {
	APP struct {
		Name    string `default:"app"`
		Version string `default:"1.0.0"`
		Addr    string `default:"8081"`
		Timeout int    `default:"5"`
	}
	Data struct {
		Database struct {
			Driver string `default:"mysql"`
			Source string
		}
		Redis struct {
			Addr         string
			Password     string
			ReadTimeout  float64 `default:"0.2"`
			WriteTimeout float64 `default:"0.2"`
		}
	}
}

func New(path string) (*Config, error) {
	// configor	支持配置文件不存在, issues: https://github.com/jinzhu/configor/issues/20
	// 为了兼容配置必须存在，需要手动判断
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	cfg := &Config{}
	// 自动重载配置
	err := configor.New(&configor.Config{AutoReload: true}).Load(cfg, path)
	return cfg, err
}
