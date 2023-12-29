package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Conf       = new(AppConfig)
	StaticPath string
)

type AppConfig struct {
	Mode           string `mapstructure:"mode"`
	Port           int    `mapstructure:"port"`
	Name           string `mapstructure:"name"`
	Version        string `mapstructure:"version"`
	StartTime      string `mapstructure:"start_time"`
	MachineID      int    `mapstructure:"machine_id"`
	StaticFilePath string `mapstructure:"static_file_path"`
	*MySQLConfig   `mapstructure:"mysql"`
	*RedisConfig   `mapstructure:"redis"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
}

func Init() error {
	viper.SetConfigFile("./conf/config.yaml") //设置读取的配置文件

	viper.WatchConfig() //监控配置文件
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("file \"config.yaml\" has been changed")
		viper.Unmarshal(&Conf)
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("readInConfig failed, err: %v", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err: %v", err))
	}
	return err
}
