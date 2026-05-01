package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type MySQLConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"-"`
	Database string `toml:"database"`
}

type LogConfig struct {
	LogPath string `toml:"logPath"`
}

type Config struct {
	MySQL MySQLConfig `toml:"mysql"`
	Log   LogConfig   `toml:"log"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if _, err := toml.DecodeFile("config/config.toml", config); err != nil {
		log.Printf("toml.DecodeFile() failed,err:%v\n", err)
		return nil, fmt.Errorf("toml.DecodeFile() failed,err:%w", err)
	}
	//开发环境模拟env，在部署环境下可以直接从环境变量中取，不需要env文件
	err := godotenv.Load()
	if err != nil {
		log.Printf("godotenv.Load() failed,err:%v\n", err)
		return nil, fmt.Errorf("godotenv.Load() failed,err:%w", err)
	}

	config.MySQL.Password = os.Getenv("MYSQL_PASSWORD")
	return config, nil
}
